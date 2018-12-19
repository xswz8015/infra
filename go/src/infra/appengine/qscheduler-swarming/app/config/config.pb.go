// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/qscheduler-swarming/app/config/config.proto

package config

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

// Config is the configuration data served by luci-config for this app.
type Config struct {
	AccessGroup string `protobuf:"bytes,1,opt,name=access_group,json=accessGroup,proto3" json:"access_group,omitempty"` // Deprecated: Do not use.
	// QuotaScheduler contains QS-specific config information.
	QuotaScheduler *QuotaScheduler `protobuf:"bytes,2,opt,name=quota_scheduler,json=quotaScheduler,proto3" json:"quota_scheduler,omitempty"`
	// Auth describes which access levels are granted to which groups.
	Auth                 *Auth    `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_07b1a5d3644599b1, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Config) GetAccessGroup() string {
	if m != nil {
		return m.AccessGroup
	}
	return ""
}

func (m *Config) GetQuotaScheduler() *QuotaScheduler {
	if m != nil {
		return m.QuotaScheduler
	}
	return nil
}

func (m *Config) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Auth struct {
	// AdminGroup is the luci-auth group controlling access to the administrative
	// endpoints of this server (the QSchedulerAdmin API).
	//
	// Members of this group also recieve QSchedulerView access.
	AdminGroup string `protobuf:"bytes,1,opt,name=admin_group,json=adminGroup,proto3" json:"admin_group,omitempty"`
	// SwarmingGroup is the luci-auth group controlling access to the scheduler
	// endpoints of this server (the swarming.ExternalScheduler API).
	SwarmingGroup string `protobuf:"bytes,2,opt,name=swarming_group,json=swarmingGroup,proto3" json:"swarming_group,omitempty"`
	// ViewGroup is the luci-auth group controlloing access to the qscheduler view
	// endpoints of the server (QSchedulerView API).
	ViewGroup            string   `protobuf:"bytes,3,opt,name=view_group,json=viewGroup,proto3" json:"view_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_07b1a5d3644599b1, []int{1}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetAdminGroup() string {
	if m != nil {
		return m.AdminGroup
	}
	return ""
}

func (m *Auth) GetSwarmingGroup() string {
	if m != nil {
		return m.SwarmingGroup
	}
	return ""
}

func (m *Auth) GetViewGroup() string {
	if m != nil {
		return m.ViewGroup
	}
	return ""
}

// QuotaScheduler contains configuration information for the QuotaScheduler app.
type QuotaScheduler struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaScheduler) Reset()         { *m = QuotaScheduler{} }
func (m *QuotaScheduler) String() string { return proto.CompactTextString(m) }
func (*QuotaScheduler) ProtoMessage()    {}
func (*QuotaScheduler) Descriptor() ([]byte, []int) {
	return fileDescriptor_07b1a5d3644599b1, []int{2}
}

func (m *QuotaScheduler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaScheduler.Unmarshal(m, b)
}
func (m *QuotaScheduler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaScheduler.Marshal(b, m, deterministic)
}
func (m *QuotaScheduler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaScheduler.Merge(m, src)
}
func (m *QuotaScheduler) XXX_Size() int {
	return xxx_messageInfo_QuotaScheduler.Size(m)
}
func (m *QuotaScheduler) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaScheduler.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaScheduler proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Config)(nil), "qscheduler.config.Config")
	proto.RegisterType((*Auth)(nil), "qscheduler.config.Auth")
	proto.RegisterType((*QuotaScheduler)(nil), "qscheduler.config.QuotaScheduler")
}

func init() {
	proto.RegisterFile("infra/appengine/qscheduler-swarming/app/config/config.proto", fileDescriptor_07b1a5d3644599b1)
}

var fileDescriptor_07b1a5d3644599b1 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0x37, 0x8a, 0x7b, 0xd5, 0xaa, 0xb9, 0xd8, 0x8b, 0x38, 0x0b, 0x83, 0x81, 0xd8,
	0x82, 0x1e, 0x3d, 0x39, 0x0f, 0x82, 0x37, 0xeb, 0xcd, 0xcb, 0x88, 0x5d, 0xd6, 0x06, 0x6c, 0xd2,
	0xa6, 0x89, 0xfb, 0x4a, 0x7e, 0x4c, 0xc9, 0x4b, 0xab, 0x0e, 0x3d, 0x25, 0xef, 0xf7, 0x7e, 0xfc,
	0xf9, 0xf3, 0xe0, 0x4e, 0xc8, 0xad, 0x66, 0x39, 0x6b, 0x5b, 0x2e, 0x2b, 0x21, 0x79, 0xde, 0xf5,
	0x65, 0xcd, 0x37, 0xf6, 0x9d, 0xeb, 0xeb, 0x7e, 0xc7, 0x74, 0x23, 0x64, 0xe5, 0xb6, 0x79, 0xa9,
	0xe4, 0x56, 0x54, 0xc3, 0x93, 0xb5, 0x5a, 0x19, 0x45, 0x4f, 0x7f, 0xe4, 0xcc, 0x2f, 0xd2, 0x4f,
	0x02, 0xe1, 0x03, 0x7e, 0xe9, 0x02, 0x0e, 0x59, 0x59, 0xf2, 0xbe, 0x5f, 0x57, 0x5a, 0xd9, 0x36,
	0x21, 0x73, 0xb2, 0x9c, 0xad, 0x82, 0x84, 0x14, 0x91, 0xe7, 0x8f, 0x0e, 0xd3, 0x27, 0x38, 0xee,
	0xac, 0x32, 0x6c, 0xfd, 0x9d, 0x95, 0x04, 0x73, 0xb2, 0x8c, 0x6e, 0x2e, 0xb3, 0x3f, 0xf1, 0xd9,
	0xb3, 0x33, 0x5f, 0x46, 0x5a, 0xc4, 0xdd, 0xde, 0x4c, 0xaf, 0x60, 0xca, 0xac, 0xa9, 0x93, 0x09,
	0x06, 0x9c, 0xfd, 0x13, 0x70, 0x6f, 0x4d, 0x5d, 0xa0, 0x94, 0x36, 0x30, 0x75, 0x13, 0xbd, 0x80,
	0x88, 0x6d, 0x1a, 0x21, 0x7f, 0xd7, 0x2c, 0x00, 0x91, 0x6f, 0xb8, 0x80, 0x78, 0x3c, 0xc5, 0xe0,
	0x04, 0xe8, 0x1c, 0x8d, 0xd4, 0x6b, 0xe7, 0x00, 0x1f, 0x82, 0xef, 0x06, 0x65, 0x82, 0xca, 0xcc,
	0x11, 0x5c, 0xa7, 0x27, 0x10, 0xef, 0xb7, 0x5f, 0x1d, 0xbc, 0x86, 0xbe, 0xd5, 0x5b, 0x88, 0xf7,
	0xbc, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xce, 0xab, 0xd6, 0x2e, 0x8e, 0x01, 0x00, 0x00,
}
