// Code generated by protoc-gen-go.
// source: job/job.proto
// DO NOT EDIT!

/*
Package job is a generated protocol buffer package.

It is generated from these files:
	job/job.proto

It has these top-level messages:
	Option
	SMCTask
	PreparePhase
	LinkingPhase
	SessionPhase
	DebugPhase
	SMCCmd
	CmdResult
	SMCResult
*/
package job

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DataOrigin int32

const (
	DataOrigin_TEMPERATURE  DataOrigin = 0
	DataOrigin_HUMIDITY     DataOrigin = 1
	DataOrigin_AMBIENT      DataOrigin = 2
	DataOrigin_AIR_PRESSURE DataOrigin = 3
	DataOrigin_PRESENCE     DataOrigin = 10
	// Dynamically assigned sensor types.
	DataOrigin_RESERVED_999            DataOrigin = 999
	DataOrigin_DYNAMIC_ASSIGNMENT_1000 DataOrigin = 1000
	// ...
	DataOrigin_DYNAMIC_ASSIGNMENT_99999 DataOrigin = 99999
	DataOrigin_RESERVED_100000          DataOrigin = 100000
)

var DataOrigin_name = map[int32]string{
	0:      "TEMPERATURE",
	1:      "HUMIDITY",
	2:      "AMBIENT",
	3:      "AIR_PRESSURE",
	10:     "PRESENCE",
	999:    "RESERVED_999",
	1000:   "DYNAMIC_ASSIGNMENT_1000",
	99999:  "DYNAMIC_ASSIGNMENT_99999",
	100000: "RESERVED_100000",
}
var DataOrigin_value = map[string]int32{
	"TEMPERATURE":              0,
	"HUMIDITY":                 1,
	"AMBIENT":                  2,
	"AIR_PRESSURE":             3,
	"PRESENCE":                 10,
	"RESERVED_999":             999,
	"DYNAMIC_ASSIGNMENT_1000":  1000,
	"DYNAMIC_ASSIGNMENT_99999": 99999,
	"RESERVED_100000":          100000,
}

func (x DataOrigin) String() string {
	return proto.EnumName(DataOrigin_name, int32(x))
}
func (DataOrigin) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Aggregator int32

const (
	Aggregator_SUM            Aggregator = 0
	Aggregator_AVG            Aggregator = 1
	Aggregator_MEDIAN         Aggregator = 2
	Aggregator_STD_DEVIATION  Aggregator = 3
	Aggregator_DBG_PINGPONG   Aggregator = 1000
	Aggregator_DBG_SET_CONFIG Aggregator = 1001
)

var Aggregator_name = map[int32]string{
	0:    "SUM",
	1:    "AVG",
	2:    "MEDIAN",
	3:    "STD_DEVIATION",
	1000: "DBG_PINGPONG",
	1001: "DBG_SET_CONFIG",
}
var Aggregator_value = map[string]int32{
	"SUM":            0,
	"AVG":            1,
	"MEDIAN":         2,
	"STD_DEVIATION":  3,
	"DBG_PINGPONG":   1000,
	"DBG_SET_CONFIG": 1001,
}

func (x Aggregator) String() string {
	return proto.EnumName(Aggregator_name, int32(x))
}
func (Aggregator) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MetadataKey int32

const (
	MetadataKey_PEERID MetadataKey = 0
)

var MetadataKey_name = map[int32]string{
	0: "PEERID",
}
var MetadataKey_value = map[string]int32{
	"PEERID": 0,
}

func (x MetadataKey) String() string {
	return proto.EnumName(MetadataKey_name, int32(x))
}
func (MetadataKey) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SMCCmd_Phase int32

const (
	SMCCmd_PREPARE SMCCmd_Phase = 0
	SMCCmd_LINK    SMCCmd_Phase = 1
	SMCCmd_SESSION SMCCmd_Phase = 2
	// ...
	SMCCmd_FINISH SMCCmd_Phase = 126
	SMCCmd_ABORT  SMCCmd_Phase = 127
)

var SMCCmd_Phase_name = map[int32]string{
	0:   "PREPARE",
	1:   "LINK",
	2:   "SESSION",
	126: "FINISH",
	127: "ABORT",
}
var SMCCmd_Phase_value = map[string]int32{
	"PREPARE": 0,
	"LINK":    1,
	"SESSION": 2,
	"FINISH":  126,
	"ABORT":   127,
}

func (x SMCCmd_Phase) String() string {
	return proto.EnumName(SMCCmd_Phase_name, int32(x))
}
func (SMCCmd_Phase) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type CmdResult_Status int32

const (
	// Class: success and info
	CmdResult_SUCCESS      CmdResult_Status = 0
	CmdResult_SUCCESS_DONE CmdResult_Status = 1
	// Class: soft or recoverable errors (32 - 63)
	CmdResult_ERR_CLASS_NORM CmdResult_Status = 32
	CmdResult_UNKNOWN_CMD    CmdResult_Status = 33
	CmdResult_DENIED         CmdResult_Status = 34
	// Class: irreversible errors on peer side (64 - 127)
	CmdResult_ERR_CLASS_FAULT CmdResult_Status = 64
	CmdResult_ABORTED         CmdResult_Status = 65
	// Class: communication errors (128 - 255)
	CmdResult_ERR_CLASS_COMM       CmdResult_Status = 128
	CmdResult_STREAM_ERR           CmdResult_Status = 129
	CmdResult_SEVERE_ERROR_CLASSES CmdResult_Status = 192
	CmdResult_ALL_ERROR_CLASSES    CmdResult_Status = 224
)

var CmdResult_Status_name = map[int32]string{
	0:   "SUCCESS",
	1:   "SUCCESS_DONE",
	32:  "ERR_CLASS_NORM",
	33:  "UNKNOWN_CMD",
	34:  "DENIED",
	64:  "ERR_CLASS_FAULT",
	65:  "ABORTED",
	128: "ERR_CLASS_COMM",
	129: "STREAM_ERR",
	192: "SEVERE_ERROR_CLASSES",
	224: "ALL_ERROR_CLASSES",
}
var CmdResult_Status_value = map[string]int32{
	"SUCCESS":              0,
	"SUCCESS_DONE":         1,
	"ERR_CLASS_NORM":       32,
	"UNKNOWN_CMD":          33,
	"DENIED":               34,
	"ERR_CLASS_FAULT":      64,
	"ABORTED":              65,
	"ERR_CLASS_COMM":       128,
	"STREAM_ERR":           129,
	"SEVERE_ERROR_CLASSES": 192,
	"ALL_ERROR_CLASSES":    224,
}

func (x CmdResult_Status) String() string {
	return proto.EnumName(CmdResult_Status_name, int32(x))
}
func (CmdResult_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type Option struct {
	// Types that are valid to be assigned to OptValue:
	//	*Option_Str
	//	*Option_Dec
	OptValue isOption_OptValue `protobuf_oneof:"optValue"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isOption_OptValue interface {
	isOption_OptValue()
}

type Option_Str struct {
	Str string `protobuf:"bytes,1,opt,name=str,oneof"`
}
type Option_Dec struct {
	Dec int32 `protobuf:"varint,2,opt,name=dec,oneof"`
}

func (*Option_Str) isOption_OptValue() {}
func (*Option_Dec) isOption_OptValue() {}

func (m *Option) GetOptValue() isOption_OptValue {
	if m != nil {
		return m.OptValue
	}
	return nil
}

func (m *Option) GetStr() string {
	if x, ok := m.GetOptValue().(*Option_Str); ok {
		return x.Str
	}
	return ""
}

func (m *Option) GetDec() int32 {
	if x, ok := m.GetOptValue().(*Option_Dec); ok {
		return x.Dec
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Option) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Option_OneofMarshaler, _Option_OneofUnmarshaler, _Option_OneofSizer, []interface{}{
		(*Option_Str)(nil),
		(*Option_Dec)(nil),
	}
}

func _Option_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Option)
	// optValue
	switch x := m.OptValue.(type) {
	case *Option_Str:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Str)
	case *Option_Dec:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Dec))
	case nil:
	default:
		return fmt.Errorf("Option.OptValue has unexpected type %T", x)
	}
	return nil
}

func _Option_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Option)
	switch tag {
	case 1: // optValue.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OptValue = &Option_Str{x}
		return true, err
	case 2: // optValue.dec
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OptValue = &Option_Dec{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Option_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Option)
	// optValue
	switch x := m.OptValue.(type) {
	case *Option_Str:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *Option_Dec:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Dec))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SMCTask struct {
	Set    string     `protobuf:"bytes,1,opt,name=set" json:"set,omitempty"`
	Source DataOrigin `protobuf:"varint,2,opt,name=source,enum=job.DataOrigin" json:"source,omitempty"`
	// (Pre)Selectors
	Aggregator      Aggregator                 `protobuf:"varint,4,opt,name=aggregator,enum=job.Aggregator" json:"aggregator,omitempty"`
	Options         map[string]*Option         `protobuf:"bytes,6,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TicketSignature string                     `protobuf:"bytes,7,opt,name=ticketSignature" json:"ticketSignature,omitempty"`
	Issued          *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=issued" json:"issued,omitempty"`
	QuerySignature  string                     `protobuf:"bytes,9,opt,name=querySignature" json:"querySignature,omitempty"`
}

func (m *SMCTask) Reset()                    { *m = SMCTask{} }
func (m *SMCTask) String() string            { return proto.CompactTextString(m) }
func (*SMCTask) ProtoMessage()               {}
func (*SMCTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SMCTask) GetOptions() map[string]*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *SMCTask) GetIssued() *google_protobuf.Timestamp {
	if m != nil {
		return m.Issued
	}
	return nil
}

type PreparePhase struct {
	Participants []*PreparePhase_Participant `protobuf:"bytes,1,rep,name=participants" json:"participants,omitempty"`
	SmcTask      *SMCTask                    `protobuf:"bytes,3,opt,name=smcTask" json:"smcTask,omitempty"`
}

func (m *PreparePhase) Reset()                    { *m = PreparePhase{} }
func (m *PreparePhase) String() string            { return proto.CompactTextString(m) }
func (*PreparePhase) ProtoMessage()               {}
func (*PreparePhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PreparePhase) GetParticipants() []*PreparePhase_Participant {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *PreparePhase) GetSmcTask() *SMCTask {
	if m != nil {
		return m.SmcTask
	}
	return nil
}

type PreparePhase_Participant struct {
	// Fixed authentication (certificate) based identity
	AuthID string `protobuf:"bytes,1,opt,name=authID" json:"authID,omitempty"`
	// Temporarily assigned peer identity for the current SMC task
	SmcPeerID int32 `protobuf:"varint,2,opt,name=smcPeerID" json:"smcPeerID,omitempty"`
	// Address and port of the service (address:port)
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *PreparePhase_Participant) Reset()                    { *m = PreparePhase_Participant{} }
func (m *PreparePhase_Participant) String() string            { return proto.CompactTextString(m) }
func (*PreparePhase_Participant) ProtoMessage()               {}
func (*PreparePhase_Participant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type LinkingPhase struct {
}

func (m *LinkingPhase) Reset()                    { *m = LinkingPhase{} }
func (m *LinkingPhase) String() string            { return proto.CompactTextString(m) }
func (*LinkingPhase) ProtoMessage()               {}
func (*LinkingPhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SessionPhase struct {
}

func (m *SessionPhase) Reset()                    { *m = SessionPhase{} }
func (m *SessionPhase) String() string            { return proto.CompactTextString(m) }
func (*SessionPhase) ProtoMessage()               {}
func (*SessionPhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DebugPhase struct {
	Ping       int32 `protobuf:"varint,1,opt,name=ping" json:"ping,omitempty"`
	MorePhases bool  `protobuf:"varint,3,opt,name=morePhases" json:"morePhases,omitempty"`
	// options configures the peers' environment. Otherwise, it is empty.
	Options map[string]*Option `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DebugPhase) Reset()                    { *m = DebugPhase{} }
func (m *DebugPhase) String() string            { return proto.CompactTextString(m) }
func (*DebugPhase) ProtoMessage()               {}
func (*DebugPhase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DebugPhase) GetOptions() map[string]*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type SMCCmd struct {
	SessionID string       `protobuf:"bytes,1,opt,name=sessionID" json:"sessionID,omitempty"`
	SmcPeerID int32        `protobuf:"varint,5,opt,name=smcPeerID" json:"smcPeerID,omitempty"`
	State     SMCCmd_Phase `protobuf:"varint,2,opt,name=state,enum=job.SMCCmd_Phase" json:"state,omitempty"`
	// Payload packet
	//
	// Types that are valid to be assigned to Payload:
	//	*SMCCmd_Prepare
	//	*SMCCmd_Link
	//	*SMCCmd_Session
	//	*SMCCmd_Debug
	Payload isSMCCmd_Payload `protobuf_oneof:"payload"`
}

func (m *SMCCmd) Reset()                    { *m = SMCCmd{} }
func (m *SMCCmd) String() string            { return proto.CompactTextString(m) }
func (*SMCCmd) ProtoMessage()               {}
func (*SMCCmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isSMCCmd_Payload interface {
	isSMCCmd_Payload()
}

type SMCCmd_Prepare struct {
	Prepare *PreparePhase `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type SMCCmd_Link struct {
	Link *LinkingPhase `protobuf:"bytes,6,opt,name=link,oneof"`
}
type SMCCmd_Session struct {
	Session *SessionPhase `protobuf:"bytes,4,opt,name=session,oneof"`
}
type SMCCmd_Debug struct {
	Debug *DebugPhase `protobuf:"bytes,9,opt,name=debug,oneof"`
}

func (*SMCCmd_Prepare) isSMCCmd_Payload() {}
func (*SMCCmd_Link) isSMCCmd_Payload()    {}
func (*SMCCmd_Session) isSMCCmd_Payload() {}
func (*SMCCmd_Debug) isSMCCmd_Payload()   {}

func (m *SMCCmd) GetPayload() isSMCCmd_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SMCCmd) GetPrepare() *PreparePhase {
	if x, ok := m.GetPayload().(*SMCCmd_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *SMCCmd) GetLink() *LinkingPhase {
	if x, ok := m.GetPayload().(*SMCCmd_Link); ok {
		return x.Link
	}
	return nil
}

func (m *SMCCmd) GetSession() *SessionPhase {
	if x, ok := m.GetPayload().(*SMCCmd_Session); ok {
		return x.Session
	}
	return nil
}

func (m *SMCCmd) GetDebug() *DebugPhase {
	if x, ok := m.GetPayload().(*SMCCmd_Debug); ok {
		return x.Debug
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SMCCmd) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SMCCmd_OneofMarshaler, _SMCCmd_OneofUnmarshaler, _SMCCmd_OneofSizer, []interface{}{
		(*SMCCmd_Prepare)(nil),
		(*SMCCmd_Link)(nil),
		(*SMCCmd_Session)(nil),
		(*SMCCmd_Debug)(nil),
	}
}

func _SMCCmd_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SMCCmd)
	// payload
	switch x := m.Payload.(type) {
	case *SMCCmd_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *SMCCmd_Link:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Link); err != nil {
			return err
		}
	case *SMCCmd_Session:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Session); err != nil {
			return err
		}
	case *SMCCmd_Debug:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Debug); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SMCCmd.Payload has unexpected type %T", x)
	}
	return nil
}

func _SMCCmd_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SMCCmd)
	switch tag {
	case 3: // payload.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PreparePhase)
		err := b.DecodeMessage(msg)
		m.Payload = &SMCCmd_Prepare{msg}
		return true, err
	case 6: // payload.link
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinkingPhase)
		err := b.DecodeMessage(msg)
		m.Payload = &SMCCmd_Link{msg}
		return true, err
	case 4: // payload.session
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SessionPhase)
		err := b.DecodeMessage(msg)
		m.Payload = &SMCCmd_Session{msg}
		return true, err
	case 9: // payload.debug
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DebugPhase)
		err := b.DecodeMessage(msg)
		m.Payload = &SMCCmd_Debug{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SMCCmd_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SMCCmd)
	// payload
	switch x := m.Payload.(type) {
	case *SMCCmd_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SMCCmd_Link:
		s := proto.Size(x.Link)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SMCCmd_Session:
		s := proto.Size(x.Session)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SMCCmd_Debug:
		s := proto.Size(x.Debug)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CmdResult struct {
	Status CmdResult_Status `protobuf:"varint,1,opt,name=status,enum=job.CmdResult_Status" json:"status,omitempty"`
	Msg    string           `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Result *SMCResult       `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	// Partly filled on receiver side (e.g. peer ID)
	Metadata map[string]string `protobuf:"bytes,7,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CmdResult) Reset()                    { *m = CmdResult{} }
func (m *CmdResult) String() string            { return proto.CompactTextString(m) }
func (*CmdResult) ProtoMessage()               {}
func (*CmdResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CmdResult) GetResult() *SMCResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CmdResult) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SMCResult struct {
	Res float64 `protobuf:"fixed64,1,opt,name=res" json:"res,omitempty"`
}

func (m *SMCResult) Reset()                    { *m = SMCResult{} }
func (m *SMCResult) String() string            { return proto.CompactTextString(m) }
func (*SMCResult) ProtoMessage()               {}
func (*SMCResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Option)(nil), "job.Option")
	proto.RegisterType((*SMCTask)(nil), "job.SMCTask")
	proto.RegisterType((*PreparePhase)(nil), "job.PreparePhase")
	proto.RegisterType((*PreparePhase_Participant)(nil), "job.PreparePhase.Participant")
	proto.RegisterType((*LinkingPhase)(nil), "job.LinkingPhase")
	proto.RegisterType((*SessionPhase)(nil), "job.SessionPhase")
	proto.RegisterType((*DebugPhase)(nil), "job.DebugPhase")
	proto.RegisterType((*SMCCmd)(nil), "job.SMCCmd")
	proto.RegisterType((*CmdResult)(nil), "job.CmdResult")
	proto.RegisterType((*SMCResult)(nil), "job.SMCResult")
	proto.RegisterEnum("job.DataOrigin", DataOrigin_name, DataOrigin_value)
	proto.RegisterEnum("job.Aggregator", Aggregator_name, Aggregator_value)
	proto.RegisterEnum("job.MetadataKey", MetadataKey_name, MetadataKey_value)
	proto.RegisterEnum("job.SMCCmd_Phase", SMCCmd_Phase_name, SMCCmd_Phase_value)
	proto.RegisterEnum("job.CmdResult_Status", CmdResult_Status_name, CmdResult_Status_value)
}

func init() { proto.RegisterFile("job/job.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0xaf, 0x93, 0x8b, 0x93, 0x4c, 0xd2, 0x9c, 0x6f, 0x69, 0x4b, 0x1a, 0xb5, 0xd0, 0x46, 0x82,
	0x56, 0x95, 0xf0, 0x95, 0x54, 0x42, 0x3d, 0xe0, 0x01, 0x27, 0xde, 0xa6, 0x6e, 0x63, 0x27, 0x5a,
	0x3b, 0x87, 0xfa, 0x14, 0xf9, 0x2e, 0xdb, 0x34, 0x5c, 0x12, 0x07, 0xdb, 0x41, 0xdc, 0x0b, 0x7f,
	0x3e, 0x40, 0x3f, 0x03, 0x88, 0x4f, 0xd2, 0x7e, 0x06, 0xde, 0x78, 0xe2, 0x0d, 0x90, 0xf8, 0x02,
	0xbc, 0x31, 0xbb, 0xb6, 0x13, 0xf7, 0xe8, 0x23, 0x91, 0x22, 0xed, 0xcc, 0xfc, 0x76, 0x76, 0xe6,
	0x37, 0xbf, 0x5d, 0xc3, 0xe5, 0xaf, 0x82, 0x93, 0x43, 0xfc, 0xeb, 0xeb, 0x30, 0x88, 0x03, 0x52,
	0xc4, 0x65, 0xeb, 0xfd, 0x59, 0x10, 0xcc, 0x16, 0xfc, 0x50, 0xba, 0x4e, 0x36, 0xcf, 0x0f, 0xe3,
	0xf9, 0x92, 0x47, 0xb1, 0xbf, 0x5c, 0x27, 0xa8, 0xf6, 0xe7, 0xa0, 0x0e, 0xd7, 0xf1, 0x3c, 0x58,
	0x11, 0x02, 0xc5, 0x28, 0x0e, 0x9b, 0xca, 0x2d, 0xe5, 0x6e, 0xf5, 0xf1, 0x25, 0x26, 0x0c, 0xe1,
	0x9b, 0xf2, 0xd3, 0x66, 0x01, 0x7d, 0x25, 0xe1, 0x43, 0xa3, 0x0b, 0x50, 0x09, 0xd6, 0xf1, 0xb1,
	0xbf, 0xd8, 0xf0, 0xf6, 0x3f, 0x05, 0x28, 0xbb, 0x76, 0xcf, 0xf3, 0xa3, 0x33, 0xa2, 0xe1, 0x7e,
	0x1e, 0x27, 0xfb, 0x99, 0x58, 0x92, 0x3b, 0xa0, 0x46, 0xc1, 0x26, 0x3c, 0xe5, 0x32, 0x41, 0xa3,
	0xb3, 0xaf, 0x8b, 0xea, 0x4c, 0x3f, 0xf6, 0x87, 0xe1, 0x7c, 0x36, 0x5f, 0xb1, 0x34, 0x4c, 0x0e,
	0x01, 0xfc, 0xd9, 0x2c, 0xe4, 0x33, 0x3f, 0x0e, 0xc2, 0xe6, 0x5e, 0x0e, 0x6c, 0x6c, 0xdd, 0x2c,
	0x07, 0x21, 0x0f, 0xa0, 0x1c, 0xc8, 0xaa, 0xa3, 0xa6, 0x7a, 0xab, 0x78, 0xb7, 0xd6, 0xb9, 0x2e,
	0xd1, 0x69, 0x29, 0x7a, 0xd2, 0x51, 0x44, 0x57, 0x71, 0x78, 0xce, 0x32, 0x24, 0xb9, 0x0b, 0xfb,
	0xf1, 0xfc, 0xf4, 0x8c, 0xc7, 0xee, 0x7c, 0xb6, 0xf2, 0xe3, 0x4d, 0xc8, 0x9b, 0x65, 0x59, 0xec,
	0x45, 0x37, 0xe9, 0x80, 0x3a, 0x8f, 0xa2, 0x0d, 0x9f, 0x36, 0x2b, 0x08, 0xa8, 0x75, 0x5a, 0x7a,
	0x42, 0xa3, 0x9e, 0xd1, 0xa8, 0x7b, 0x19, 0x8d, 0x2c, 0x45, 0x92, 0x0f, 0xa1, 0xf1, 0xf5, 0x86,
	0x87, 0xe7, 0xbb, 0xe4, 0x55, 0x99, 0xfc, 0x82, 0xb7, 0xd5, 0x87, 0x7a, 0xbe, 0x3c, 0x41, 0xdb,
	0x19, 0x3f, 0xcf, 0x68, 0xc3, 0x25, 0xb9, 0x0d, 0xa5, 0x6f, 0x04, 0xbb, 0x92, 0xb5, 0x5a, 0xa7,
	0x26, 0x5b, 0x4b, 0xf6, 0xb0, 0x24, 0xf2, 0x69, 0xe1, 0xa1, 0xd2, 0xfe, 0x4d, 0x81, 0xfa, 0x28,
	0xe4, 0x6b, 0x3f, 0xe4, 0xa3, 0x17, 0x7e, 0xc4, 0x89, 0x01, 0x75, 0x34, 0xb0, 0x97, 0xf9, 0xda,
	0x5f, 0xc5, 0x11, 0xa6, 0x14, 0xcc, 0xdc, 0x94, 0xdb, 0xf3, 0x40, 0x7d, 0xb4, 0x43, 0xb1, 0x37,
	0xb6, 0x60, 0x13, 0xe5, 0x68, 0x79, 0x2a, 0x38, 0x6c, 0x16, 0xe5, 0xe1, 0xf5, 0x3c, 0xaf, 0x2c,
	0x0b, 0xb6, 0x26, 0x50, 0xcb, 0x25, 0x21, 0xd7, 0x40, 0xf5, 0x37, 0xf1, 0x0b, 0xcb, 0x4c, 0xdb,
	0x48, 0x2d, 0x72, 0x03, 0xaa, 0xb8, 0x63, 0xc4, 0x79, 0x88, 0x21, 0x29, 0x22, 0xb6, 0x73, 0x90,
	0x16, 0x54, 0xf8, 0x6a, 0xba, 0x0e, 0xe6, 0xab, 0x58, 0x9e, 0x56, 0x65, 0x5b, 0xbb, 0xdd, 0x80,
	0xfa, 0x60, 0xbe, 0x3a, 0x9b, 0xaf, 0x66, 0xb2, 0x64, 0x61, 0xbb, 0x3c, 0x8a, 0x90, 0x82, 0xc4,
	0x7e, 0xad, 0x00, 0x98, 0xfc, 0x64, 0x93, 0x84, 0x51, 0xa7, 0x7b, 0x6b, 0xc4, 0xca, 0xe3, 0x4b,
	0x4c, 0xae, 0xc9, 0x7b, 0x00, 0xcb, 0x20, 0x6d, 0x39, 0x92, 0x07, 0x54, 0x58, 0xce, 0x43, 0x3e,
	0xd9, 0x69, 0xa8, 0x20, 0x99, 0xba, 0x91, 0xc8, 0x73, 0x9b, 0xf5, 0xed, 0x32, 0xfa, 0xff, 0x06,
	0xf8, 0x77, 0x01, 0x54, 0x64, 0xb6, 0xb7, 0x9c, 0x4a, 0xa2, 0x92, 0xf6, 0xb6, 0x1c, 0xee, 0x1c,
	0x6f, 0xd2, 0x58, 0xba, 0x48, 0xe3, 0x1d, 0x28, 0xa1, 0x12, 0xe3, 0xec, 0x92, 0x1d, 0x64, 0x13,
	0xc3, 0xbc, 0xba, 0xec, 0x83, 0x25, 0x71, 0xf2, 0x11, 0x94, 0xd7, 0x89, 0x0c, 0xd2, 0xe1, 0x1e,
	0xfc, 0x47, 0x1a, 0x78, 0xc7, 0x33, 0x0c, 0xe6, 0xdd, 0x5b, 0xe0, 0x08, 0xf0, 0x82, 0xed, 0xb0,
	0xf9, 0x99, 0x20, 0x56, 0x02, 0x44, 0xde, 0xb4, 0x56, 0x79, 0x75, 0x33, 0x6c, 0x7e, 0x5e, 0x22,
	0x6f, 0x8a, 0x11, 0xf5, 0x4e, 0x05, 0xc7, 0xf2, 0x7e, 0xd4, 0xb2, 0x47, 0x61, 0xcb, 0x3a, 0x42,
	0x93, 0x78, 0xbb, 0x0b, 0xa5, 0x64, 0xba, 0x35, 0x28, 0x8f, 0x18, 0x1d, 0x19, 0x8c, 0x6a, 0x97,
	0x48, 0x05, 0xf6, 0x06, 0x96, 0xf3, 0x54, 0x53, 0x84, 0xdb, 0xa5, 0xae, 0x6b, 0x0d, 0x1d, 0xad,
	0x40, 0x00, 0xd4, 0x47, 0x96, 0x63, 0xb9, 0x8f, 0xb5, 0xef, 0x48, 0x15, 0x4a, 0x46, 0x77, 0xc8,
	0x3c, 0xed, 0xfb, 0x6e, 0x15, 0x7b, 0xf6, 0xcf, 0x17, 0x81, 0x3f, 0x6d, 0xbf, 0x2e, 0x42, 0x15,
	0x39, 0x61, 0x3c, 0xda, 0x2c, 0x62, 0x2c, 0x5a, 0x15, 0xac, 0x6c, 0x22, 0x49, 0x77, 0xa3, 0x73,
	0x55, 0x96, 0xb1, 0x8d, 0xeb, 0xae, 0x0c, 0xb2, 0x14, 0x24, 0x86, 0xbc, 0x8c, 0x66, 0x92, 0x62,
	0x1c, 0x32, 0x2e, 0xf1, 0xaa, 0xa8, 0xa1, 0x84, 0xa6, 0x64, 0x36, 0x32, 0xde, 0x93, 0x04, 0x2c,
	0x8d, 0x92, 0x87, 0x50, 0x59, 0xf2, 0xd8, 0x9f, 0xe2, 0xab, 0x87, 0xcf, 0xcd, 0x4e, 0x67, 0xbb,
	0xa3, 0xec, 0x34, 0x9c, 0xe8, 0x6c, 0x8b, 0x6e, 0x7d, 0x06, 0x97, 0xdf, 0x08, 0xbd, 0x45, 0x69,
	0x57, 0xf2, 0x4a, 0xab, 0xe6, 0xc5, 0xf5, 0xab, 0x82, 0xe2, 0x4a, 0x6a, 0x17, 0x3c, 0x8d, 0x7b,
	0x3d, 0xa4, 0x0a, 0xe9, 0xd3, 0xf0, 0x22, 0x25, 0xc6, 0xc4, 0x1c, 0x3a, 0x14, 0x69, 0x24, 0xd0,
	0xa0, 0x8c, 0x4d, 0x7a, 0x03, 0x03, 0x7d, 0xce, 0x90, 0xd9, 0xda, 0x2d, 0xb2, 0x0f, 0xb5, 0xb1,
	0xf3, 0xd4, 0x19, 0x7e, 0xe9, 0x4c, 0x7a, 0xb6, 0xa9, 0xdd, 0x16, 0xf4, 0x9a, 0xd4, 0xb1, 0xa8,
	0xa9, 0xb5, 0xc9, 0x3b, 0xb0, 0xbf, 0xdb, 0xf0, 0xc8, 0x18, 0x0f, 0x3c, 0xed, 0x0b, 0x71, 0x88,
	0xe4, 0x1c, 0x11, 0x06, 0x22, 0x72, 0x29, 0x7b, 0x43, 0xdb, 0xd6, 0x7e, 0x50, 0x30, 0x27, 0xb8,
	0x1e, 0xa3, 0x86, 0x3d, 0xc1, 0x98, 0xf6, 0xa3, 0x42, 0xae, 0xc3, 0x15, 0x97, 0x1e, 0x53, 0x46,
	0x85, 0x63, 0x98, 0xc2, 0xa9, 0xab, 0xbd, 0x52, 0xf0, 0x41, 0x39, 0x30, 0x06, 0x83, 0x0b, 0xfe,
	0xdf, 0x95, 0xf6, 0x4d, 0xa8, 0x6e, 0x19, 0x16, 0x74, 0x20, 0xc7, 0x92, 0x0e, 0x85, 0x89, 0xe5,
	0xbd, 0x57, 0xe2, 0x55, 0xd8, 0x7e, 0x5e, 0x44, 0x17, 0x1e, 0xb5, 0x47, 0x94, 0x19, 0xde, 0x58,
	0x6a, 0xa7, 0x0e, 0x95, 0xc7, 0x63, 0xdb, 0x32, 0x2d, 0xef, 0x59, 0xa2, 0x1f, 0xc3, 0xee, 0x5a,
	0xd4, 0xf1, 0x50, 0x3f, 0xc8, 0x8b, 0x61, 0xb1, 0x09, 0xea, 0xcc, 0x75, 0x05, 0xb8, 0x28, 0xc0,
	0xc2, 0xa2, 0x4e, 0x8f, 0x6a, 0x40, 0x0e, 0xa0, 0x2e, 0x0c, 0x76, 0x4c, 0xcd, 0xc9, 0xd1, 0xd1,
	0x91, 0xf6, 0x47, 0x19, 0xaf, 0xe5, 0xbb, 0xe6, 0x33, 0xc7, 0xb0, 0xad, 0xde, 0x04, 0x0b, 0xb4,
	0xfa, 0x8e, 0x8d, 0xa9, 0x26, 0x1f, 0xdf, 0xbf, 0x7f, 0x5f, 0xfb, 0xb3, 0x8c, 0xcf, 0x4f, 0xf3,
	0x2d, 0xd1, 0x23, 0xf1, 0xd3, 0x7e, 0x7a, 0xa9, 0x92, 0xab, 0xb0, 0xbf, 0x4d, 0x28, 0xf6, 0xe0,
	0xae, 0x9f, 0x5f, 0xaa, 0xf7, 0x38, 0xc0, 0xee, 0x9b, 0x47, 0xca, 0x50, 0x74, 0xc7, 0x36, 0x56,
	0x8e, 0x0b, 0xe3, 0xb8, 0x8f, 0x45, 0xe3, 0x20, 0x6c, 0x6a, 0x5a, 0x86, 0xd0, 0xfc, 0x01, 0x5c,
	0x76, 0x3d, 0x73, 0x62, 0xd2, 0x63, 0xcb, 0xf0, 0xc4, 0x35, 0x28, 0x8a, 0x32, 0xcd, 0x6e, 0x7f,
	0x32, 0xb2, 0x9c, 0xfe, 0x68, 0xe8, 0xf4, 0x45, 0x21, 0x38, 0x0c, 0xe1, 0x72, 0xa9, 0x87, 0xa3,
	0x70, 0x1e, 0x59, 0x7d, 0xed, 0xaf, 0xf2, 0xbd, 0xeb, 0x50, 0xcb, 0xb4, 0xf5, 0x14, 0x75, 0x84,
	0x59, 0x47, 0x94, 0x32, 0xcb, 0xd4, 0x2e, 0x75, 0x3f, 0x80, 0x6b, 0x53, 0xae, 0xc7, 0x9b, 0xa5,
	0xfe, 0x7c, 0xc1, 0xbf, 0xc5, 0x77, 0x46, 0x17, 0xff, 0x70, 0x7d, 0xda, 0x2d, 0x3e, 0x09, 0x4e,
	0x46, 0xca, 0x2f, 0x85, 0xe2, 0x93, 0x61, 0xf7, 0x44, 0x95, 0xdf, 0xc2, 0x07, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x38, 0x99, 0xde, 0x44, 0x76, 0x08, 0x00, 0x00,
}
