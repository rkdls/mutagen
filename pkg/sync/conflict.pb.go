// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/sync/conflict.proto

package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Conflict struct {
	AlphaChanges         []*Change `protobuf:"bytes,1,rep,name=alphaChanges,proto3" json:"alphaChanges,omitempty"`
	BetaChanges          []*Change `protobuf:"bytes,2,rep,name=betaChanges,proto3" json:"betaChanges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Conflict) Reset()         { *m = Conflict{} }
func (m *Conflict) String() string { return proto.CompactTextString(m) }
func (*Conflict) ProtoMessage()    {}
func (*Conflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_conflict_68ba2f18eecd7183, []int{0}
}
func (m *Conflict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conflict.Unmarshal(m, b)
}
func (m *Conflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conflict.Marshal(b, m, deterministic)
}
func (dst *Conflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conflict.Merge(dst, src)
}
func (m *Conflict) XXX_Size() int {
	return xxx_messageInfo_Conflict.Size(m)
}
func (m *Conflict) XXX_DiscardUnknown() {
	xxx_messageInfo_Conflict.DiscardUnknown(m)
}

var xxx_messageInfo_Conflict proto.InternalMessageInfo

func (m *Conflict) GetAlphaChanges() []*Change {
	if m != nil {
		return m.AlphaChanges
	}
	return nil
}

func (m *Conflict) GetBetaChanges() []*Change {
	if m != nil {
		return m.BetaChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*Conflict)(nil), "sync.Conflict")
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/sync/conflict.proto", fileDescriptor_conflict_68ba2f18eecd7183)
}

var fileDescriptor_conflict_68ba2f18eecd7183 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7,
	0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0xae, 0xcc, 0x4b,
	0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0xc9, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x01, 0x09, 0x4a, 0x19, 0x12, 0xa7, 0x35, 0x23, 0x31, 0x2f, 0x3d, 0x15, 0xa2, 0x51, 0x29,
	0x87, 0x8b, 0xc3, 0x19, 0x6a, 0x94, 0x90, 0x01, 0x17, 0x4f, 0x62, 0x4e, 0x41, 0x46, 0xa2, 0x33,
	0x58, 0x41, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x8f, 0x1e, 0x48, 0x97, 0x1e, 0x44,
	0x30, 0x08, 0x45, 0x85, 0x90, 0x1e, 0x17, 0x77, 0x52, 0x6a, 0x09, 0x5c, 0x03, 0x13, 0x16, 0x0d,
	0xc8, 0x0a, 0x92, 0xd8, 0xc0, 0x96, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x8d, 0x85,
	0x14, 0xe4, 0x00, 0x00, 0x00,
}
