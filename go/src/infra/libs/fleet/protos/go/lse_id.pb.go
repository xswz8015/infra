// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lse_id.proto

package fleet

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

type LabSetupEnvID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabSetupEnvID) Reset()         { *m = LabSetupEnvID{} }
func (m *LabSetupEnvID) String() string { return proto.CompactTextString(m) }
func (*LabSetupEnvID) ProtoMessage()    {}
func (*LabSetupEnvID) Descriptor() ([]byte, []int) {
	return fileDescriptor_93308ec8ab98bd13, []int{0}
}

func (m *LabSetupEnvID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabSetupEnvID.Unmarshal(m, b)
}
func (m *LabSetupEnvID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabSetupEnvID.Marshal(b, m, deterministic)
}
func (m *LabSetupEnvID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabSetupEnvID.Merge(m, src)
}
func (m *LabSetupEnvID) XXX_Size() int {
	return xxx_messageInfo_LabSetupEnvID.Size(m)
}
func (m *LabSetupEnvID) XXX_DiscardUnknown() {
	xxx_messageInfo_LabSetupEnvID.DiscardUnknown(m)
}

var xxx_messageInfo_LabSetupEnvID proto.InternalMessageInfo

func (m *LabSetupEnvID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*LabSetupEnvID)(nil), "fleet.LabSetupEnvID")
}

func init() { proto.RegisterFile("lse_id.proto", fileDescriptor_93308ec8ab98bd13) }

var fileDescriptor_93308ec8ab98bd13 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x29, 0x4e, 0x8d,
	0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xcb, 0x49, 0x4d, 0x2d, 0x51,
	0x52, 0xe5, 0xe2, 0xf5, 0x49, 0x4c, 0x0a, 0x4e, 0x2d, 0x29, 0x2d, 0x70, 0xcd, 0x2b, 0xf3, 0x74,
	0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0x70, 0x9c, 0x38, 0xa3, 0xd8, 0xf5, 0xac, 0xc1, 0x3a, 0x92, 0xd8, 0xc0, 0xfa, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x23, 0x0d, 0xd6, 0x4f, 0x00, 0x00, 0x00,
}
