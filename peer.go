package main

import (
	"math"
	"time"

	gtypeAny "github.com/golang/protobuf/ptypes/any"
	"github.com/grandcat/flexsmc/logs"
	"github.com/grandcat/flexsmc/modules"
	proto "github.com/grandcat/flexsmc/proto"
	"github.com/grandcat/flexsmc/smc"
	"github.com/grandcat/srpc/client"
	"github.com/grandcat/srpc/pairing"
	"golang.org/x/net/context"
)

const RetryConnection = 5

type RunState uint8

const (
	Discovery RunState = iota
	Pairing
	Resolving
	Connecting
	Operating
)

type Peer struct {
	srpc    *client.Client
	smcConn smc.Connector
	opts    PeerOptions

	state     RunState
	connRetry int
	faults    int

	gclient   proto.GatewayClient
	modInfo   modules.ModuleContext
	modCancel context.CancelFunc
}

func NewPeer(opts PeerOptions) *Peer {
	srpcOpts := append(opts.SRpcOpts, client.TLSKeyFile(*certFile, *keyFile))

	smcConn := opts.smcBackend
	if smcConn == nil {
		smcConn = smc.DefaultSMCConnector("")
	}

	return &Peer{
		srpc:    client.NewClient(srpcOpts...),
		smcConn: smcConn,
		opts:    opts,
	}
}

func (p *Peer) Init() {
	if err := p.srpc.GetPeerCerts().LoadFromPath("peer1/"); err != nil {
		logs.Errorln(err)
	}
}

func (p *Peer) Operate() {
	for {
		logs.I.Infof("Entering %+v", p.state)
		switch p.state {
		case Discovery:
			// XXX: no discovery yet. Just go to next state
			time.Sleep(time.Second * 1)
			p.state = p.discover()

		case Pairing:
			p.state = p.startPairing()

		case Resolving:
			p.state = p.prepare()

		case Connecting:
			// Start services if connection is available:
			// - Health report
			// - SMC advisor
			p.state = p.startService()

		case Operating:
			p.state = p.watchService()
		}
	}
}

func (p *Peer) discover() (next RunState) {
	// If not known, initiate pairing
	// knownGW := p.srpc.GetPeerCerts().ActivePeerCertificates("gw4242.flexsmc.local")
	knownGW := 0
	if knownGW == 0 {
		next = Pairing

	} else {
		next = Resolving
	}
	return
}

func (p *Peer) startPairing() (next RunState) {
	const peerID = "gw4242.flexsmc.local"
	// 2. Initiate pairing if it is an unknown identity (if desired)
	// knownGW := p.srpc.GetPeerCerts().ActivePeerCertificates(peerID)
	logs.VV.Infoln("Pairing active?", p.opts.UsePairing)
	if p.opts.UsePairing {
		logs.I.Infoln("Start pairing...")
		// XXX: assume we want to pair with this gateway (e.g. matching properties, instructed by admin, ...)
		ccp, err := p.srpc.DialUnsecure(peerID)
		if err != nil {
			logs.Errorf("Could not initiate pairing to GW node %s: %v", peerID, err)
			return
		}
		logs.I.Infoln("DialUnsecured done.")
		mPairing := pairing.NewClientApproval(p.srpc.GetPeerCerts(), ccp)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		gwIdentity, err := mPairing.StartPairing(ctx, &gtypeAny.Any{"flexsmc/peerinfo", []byte(p.opts.NodeInfo)})
		if err != nil {
			logs.Errorf("Pairing with %s failed: %v", peerID, err)
			return
		}
		// Pairing commissioning
		logs.I.Infoln("GWIdentity:", gwIdentity.Fingerprint(), "Info:", gwIdentity.Details())
		// XXX: accept GW without out-of-band verification for now
		gwIdentity.Accept()
		p.srpc.GetPeerCerts().StoreToPath("peer1/")
		// Wait for server to accept our pairing request
		status := mPairing.AwaitPairingResult(ctx)
		if r, ok := <-status; ok {
			logs.I.Infoln("Pairing: peer responded with", r)
		} else {
			logs.I.Infoln("Pairing aborted by peer")
		}
	}

	next = Resolving
	return
}

func (p *Peer) prepare() (next RunState) {
	const peerID = "gw4242.flexsmc.local"
	// Join the SMC network.
	cc, err := p.srpc.Dial(peerID)
	if err != nil {
		logs.Warningf("Could not resolve or dial GW node %s: %v", peerID, err)
		next = Discovery
		return
	}
	p.gclient = proto.NewGatewayClient(cc)
	next = Connecting
	return
}

func (p *Peer) startService() (next RunState) {
	logs.I.Infoln("Services:")

	var modCtx context.Context
	modCtx, p.modCancel = context.WithCancel(context.Background())
	p.modInfo = modules.NewModuleContext(modCtx, p.gclient)

	// Health report
	healthMod := modules.NewHealthReporter(p.modInfo, time.Second*10)
	if err := healthMod.Ping(); err != nil {
		p.connRetry++
		logs.I.Infoln("[ ] Health ping")
		// Watch failed connection attempts, retry or abort if unavailable
		switch {
		case 0 < p.connRetry && p.connRetry < RetryConnection:
			timeout := math.Pow(2, float64(p.connRetry))
			logs.V.Infof("Retrying in %f seconds...", timeout)
			time.Sleep(time.Second * time.Duration(timeout))
			next = Connecting

		case RetryConnection <= p.connRetry:
			// Cleanup for new try
			p.connRetry = 0
			p.srpc.TearDown()

			next = Discovery
		}
		return
	}
	p.connRetry = 0
	healthMod.Start()
	logs.I.Infoln("[x] Health ping")

	time.Sleep(time.Millisecond * 500)

	// Start SMC advisor to bridge instructions sent by network to a SMC backend.
	smcAdvisor := modules.NewSMCAdvisor(p.modInfo, p.smcConn)
	// SpawnListener ends at the moment when finishing a SMC session.
	// Need to respawn on time.
	logs.I.Infoln("[x] SMC advisor")
	smcAdvisor.Start()

	next = Operating
	return
}

func (p *Peer) watchService() (next RunState) {
	if modErr, ok := <-p.modInfo.Faults(); ok {
		logs.Warningf("A module cannot recover from: %v", modErr.Error())
	}
	// Cancel all running modules for a fresh connection.
	p.faults++
	p.modCancel()
	p.modInfo.WaitAll()
	logs.I.Infof("!!! %d faults by modules so far. Restarting all modules with new connection.", p.faults)
	// p.srpc.TearDown()

	next = Connecting
	return
}

func (p *Peer) Run() {
	p.Operate()
	p.srpc.TearDown()
}
