package modules

import (
	"time"

	"github.com/grandcat/flexsmc/logs"
	proto "github.com/grandcat/flexsmc/proto"
)

type HealthReporter struct {
	ModuleContext
	PingInterval time.Duration
}

func NewHealthReporter(modInfo ModuleContext, pingInterval time.Duration) *HealthReporter {
	return &HealthReporter{
		ModuleContext: modInfo,
		PingInterval:  pingInterval,
	}
}

func (h *HealthReporter) Start() {
	h.ActiveMods.Add(1)
	go h.reportloop()
}

func (h *HealthReporter) reportloop() {
	defer h.ActiveMods.Done()

	ticker := time.NewTicker(h.PingInterval)
	defer ticker.Stop()

	for {
		// Ping
		if err := h.Ping(); err != nil {
			return
		}
		// Schedule next or abort
		select {
		case <-ticker.C:
			// Do nothing and ping again on next round :-)

		case <-h.context.Done():
			// Abort by context
			logs.I.Infof("Health reporter aborted: %v", h.context.Err())
			return
		}
	}
}

func (h *HealthReporter) Ping() error {
	resp, err := h.GWConn.Ping(h.context, &proto.SMCInfo{12345})
	if err != nil {
		logs.I.Infof("Could not ping GW: %v", err)
		return err
	}
	logs.V.Infoln("Ping resp:", resp.Status)
	return nil
}
