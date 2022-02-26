// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeState int32

const (
	NodeState_down NodeState = 0
	NodeState_up   NodeState = 1
)

var NodeState_name = map[int32]string{
	0: "down",
	1: "up",
}

var NodeState_value = map[string]int32{
	"down": 0,
	"up":   1,
}

func (x NodeState) String() string {
	return proto.EnumName(NodeState_name, int32(x))
}

func (NodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

type AgentState int32

const (
	AgentState_healthy   AgentState = 0
	AgentState_unhealthy AgentState = 1
)

var AgentState_name = map[int32]string{
	0: "healthy",
	1: "unhealthy",
}

var AgentState_value = map[string]int32{
	"healthy":   0,
	"unhealthy": 1,
}

func (x AgentState) String() string {
	return proto.EnumName(AgentState_name, int32(x))
}

func (AgentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

type MessageType int32

const (
	MessageType_metric MessageType = 0
	MessageType_event  MessageType = 1
)

var MessageType_name = map[int32]string{
	0: "metric",
	1: "event",
}

var MessageType_value = map[string]int32{
	"metric": 0,
	"event":  1,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

type Message struct {
	Timestamp            int64       `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type                 MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=metrika.MessageType" json:"type,omitempty"`
	NodeState            NodeState   `protobuf:"varint,3,opt,name=nodeState,proto3,enum=metrika.NodeState" json:"nodeState,omitempty"`
	AgentState           AgentState  `protobuf:"varint,4,opt,name=agentState,proto3,enum=metrika.AgentState" json:"agentState,omitempty"`
	Name                 string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Body                 []byte      `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_metric
}

func (m *Message) GetNodeState() NodeState {
	if m != nil {
		return m.NodeState
	}
	return NodeState_down
}

func (m *Message) GetAgentState() AgentState {
	if m != nil {
		return m.AgentState
	}
	return AgentState_healthy
}

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Event struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Values               []byte   `protobuf:"bytes,3,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Event) GetValues() []byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type PlatformMessage struct {
	Data                 []*Message `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	AgentUUID            string     `protobuf:"bytes,2,opt,name=agentUUID,proto3" json:"agentUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PlatformMessage) Reset()         { *m = PlatformMessage{} }
func (m *PlatformMessage) String() string { return proto.CompactTextString(m) }
func (*PlatformMessage) ProtoMessage()    {}
func (*PlatformMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *PlatformMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformMessage.Unmarshal(m, b)
}
func (m *PlatformMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformMessage.Marshal(b, m, deterministic)
}
func (m *PlatformMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformMessage.Merge(m, src)
}
func (m *PlatformMessage) XXX_Size() int {
	return xxx_messageInfo_PlatformMessage.Size(m)
}
func (m *PlatformMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformMessage proto.InternalMessageInfo

func (m *PlatformMessage) GetData() []*Message {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PlatformMessage) GetAgentUUID() string {
	if m != nil {
		return m.AgentUUID
	}
	return ""
}

type PlatformResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformResponse) Reset()         { *m = PlatformResponse{} }
func (m *PlatformResponse) String() string { return proto.CompactTextString(m) }
func (*PlatformResponse) ProtoMessage()    {}
func (*PlatformResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *PlatformResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformResponse.Unmarshal(m, b)
}
func (m *PlatformResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformResponse.Marshal(b, m, deterministic)
}
func (m *PlatformResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformResponse.Merge(m, src)
}
func (m *PlatformResponse) XXX_Size() int {
	return xxx_messageInfo_PlatformResponse.Size(m)
}
func (m *PlatformResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformResponse proto.InternalMessageInfo

func (m *PlatformResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("metrika.NodeState", NodeState_name, NodeState_value)
	proto.RegisterEnum("metrika.AgentState", AgentState_name, AgentState_value)
	proto.RegisterEnum("metrika.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "metrika.Message")
	proto.RegisterType((*Event)(nil), "metrika.Event")
	proto.RegisterType((*PlatformMessage)(nil), "metrika.PlatformMessage")
	proto.RegisterType((*PlatformResponse)(nil), "metrika.PlatformResponse")
}

func init() {
	proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77)
}

var fileDescriptor_56ede974c0020f77 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0xaf, 0xd3, 0x30,
	0x10, 0xac, 0x5f, 0x93, 0xb4, 0xd9, 0x3c, 0x20, 0x5a, 0x10, 0x0a, 0x08, 0xa4, 0x28, 0x7a, 0x87,
	0xa8, 0x87, 0x50, 0xb5, 0x47, 0x0e, 0x08, 0x04, 0x02, 0x0e, 0x20, 0x64, 0xda, 0x0b, 0x37, 0xb7,
	0x59, 0xda, 0x8a, 0xc4, 0x8e, 0x62, 0xb7, 0x28, 0xff, 0x96, 0x9f, 0x82, 0x6a, 0xf2, 0x51, 0xca,
	0x81, 0xdb, 0x7a, 0x67, 0xc6, 0x99, 0x19, 0x07, 0x02, 0xb1, 0x23, 0x69, 0xb2, 0xaa, 0x56, 0x46,
	0xe1, 0xa4, 0x24, 0x53, 0x1f, 0x7e, 0x88, 0xe4, 0x17, 0x83, 0xc9, 0x27, 0xd2, 0x5a, 0xec, 0x08,
	0x9f, 0x81, 0x6f, 0x0e, 0x25, 0x69, 0x23, 0xca, 0x2a, 0x62, 0x31, 0x4b, 0xc7, 0x7c, 0x58, 0x60,
	0x0a, 0x8e, 0x69, 0x2a, 0x8a, 0x6e, 0x62, 0x96, 0xde, 0x5f, 0x3c, 0xca, 0xda, 0x1b, 0xb2, 0x56,
	0xbd, 0x6a, 0x2a, 0xe2, 0x96, 0x81, 0x73, 0xf0, 0xa5, 0xca, 0xe9, 0xab, 0x11, 0x86, 0xa2, 0xb1,
	0xa5, 0x63, 0x4f, 0xff, 0xdc, 0x21, 0x7c, 0x20, 0xe1, 0x12, 0xc0, 0xba, 0xfb, 0x23, 0x71, 0xac,
	0xe4, 0x61, 0x2f, 0x79, 0xdd, 0x43, 0xfc, 0x82, 0x86, 0x08, 0x8e, 0x14, 0x25, 0x45, 0x6e, 0xcc,
	0x52, 0x9f, 0xdb, 0xf9, 0xbc, 0xdb, 0xa8, 0xbc, 0x89, 0xbc, 0x98, 0xa5, 0xb7, 0xdc, 0xce, 0xc9,
	0x7b, 0x70, 0xdf, 0x9d, 0x48, 0x9a, 0x5e, 0xc0, 0xfe, 0x16, 0xe4, 0xa4, 0xb7, 0x36, 0x95, 0xcf,
	0xed, 0x8c, 0x8f, 0xc1, 0x3b, 0x89, 0xe2, 0x48, 0xda, 0x9a, 0xbf, 0xe5, 0xed, 0x29, 0x59, 0xc3,
	0x83, 0x2f, 0x85, 0x30, 0xdf, 0x55, 0x5d, 0x76, 0x95, 0xdd, 0x81, 0x93, 0x0b, 0x23, 0x22, 0x16,
	0x8f, 0xd3, 0x60, 0x11, 0x5e, 0x97, 0xc2, 0x2d, 0x7a, 0x2e, 0xd6, 0xfa, 0x5e, 0xaf, 0x3f, 0xbe,
	0x6d, 0xbf, 0x34, 0x2c, 0x92, 0x39, 0x84, 0xdd, 0xb5, 0x9c, 0x74, 0xa5, 0xa4, 0xfe, 0xcf, 0x53,
	0xcc, 0x9e, 0x83, 0xdf, 0xd7, 0x88, 0x53, 0x70, 0x72, 0xf5, 0x53, 0x86, 0x23, 0xf4, 0xe0, 0xe6,
	0x58, 0x85, 0x6c, 0x96, 0x02, 0x0c, 0x95, 0x61, 0x00, 0x93, 0x3d, 0x89, 0xc2, 0xec, 0x9b, 0x70,
	0x84, 0xf7, 0xc0, 0x3f, 0xca, 0xee, 0xc8, 0x66, 0x77, 0x10, 0x5c, 0x3c, 0x1f, 0x02, 0x78, 0x36,
	0xc0, 0x36, 0x1c, 0xa1, 0x0f, 0x2e, 0x9d, 0x5b, 0x0b, 0xd9, 0xe2, 0x03, 0xb8, 0xd6, 0x2d, 0xbe,
	0x82, 0xe9, 0xaa, 0x16, 0x52, 0x97, 0x07, 0x83, 0x51, 0x9f, 0xf5, 0xaa, 0x93, 0xa7, 0x4f, 0xfe,
	0x41, 0xba, 0x58, 0x6f, 0xe0, 0xdb, 0x34, 0x7b, 0xf1, 0xb2, 0x54, 0x39, 0x15, 0x1b, 0xcf, 0xfe,
	0x89, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xcc, 0xad, 0x43, 0x98, 0x02, 0x00, 0x00,
}