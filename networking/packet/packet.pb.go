// Code generated by protoc-gen-go.
// source: packet.proto
// DO NOT EDIT!

package packet

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Packet struct {
	EnsureArrival    *uint64               `protobuf:"varint,1,opt,name=ensureArrival" json:"ensureArrival,omitempty"`
	ArrivalNotice    *Packet_ArrivalNotice `protobuf:"bytes,2,opt,name=arrivalNotice" json:"arrivalNotice,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}

func (m *Packet) GetEnsureArrival() uint64 {
	if m != nil && m.EnsureArrival != nil {
		return *m.EnsureArrival
	}
	return 0
}

func (m *Packet) GetArrivalNotice() *Packet_ArrivalNotice {
	if m != nil {
		return m.ArrivalNotice
	}
	return nil
}

type Packet_ArrivalNotice struct {
	PacketID         *uint64 `protobuf:"varint,1,req,name=packetID" json:"packetID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Packet_ArrivalNotice) Reset()         { *m = Packet_ArrivalNotice{} }
func (m *Packet_ArrivalNotice) String() string { return proto.CompactTextString(m) }
func (*Packet_ArrivalNotice) ProtoMessage()    {}

func (m *Packet_ArrivalNotice) GetPacketID() uint64 {
	if m != nil && m.PacketID != nil {
		return *m.PacketID
	}
	return 0
}

func init() {
}
