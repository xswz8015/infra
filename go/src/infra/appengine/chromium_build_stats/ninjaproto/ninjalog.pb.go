// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ninjalog.proto

package ninjaproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// OS.
type NinjaTask_OS int32

const (
	NinjaTask_UNKNOWN NinjaTask_OS = 0
	NinjaTask_WIN     NinjaTask_OS = 1
	NinjaTask_LINUX   NinjaTask_OS = 2
	NinjaTask_MAC     NinjaTask_OS = 3
)

var NinjaTask_OS_name = map[int32]string{
	0: "UNKNOWN",
	1: "WIN",
	2: "LINUX",
	3: "MAC",
}

var NinjaTask_OS_value = map[string]int32{
	"UNKNOWN": 0,
	"WIN":     1,
	"LINUX":   2,
	"MAC":     3,
}

func (x NinjaTask_OS) String() string {
	return proto.EnumName(NinjaTask_OS_name, int32(x))
}

func (NinjaTask_OS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11621a7b7e2b33ea, []int{0, 0}
}

// NinjaTask message for the log uploaded from chromium developers or buildbot.
// Due to row size limit (1MB) of BigQuery streaming insert, this message
// corresponds to the one task of ninja_log.
// NEXT ID TO USE: 14
type NinjaTask struct {
	// Identifier of build used in buildbucket api v2 (go/buildbucket-api-v2)
	// Or some random number representing an invocation of build.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// Build targets passed to ninja.
	Targets []string `protobuf:"bytes,8,rep,name=targets,proto3" json:"targets,omitempty"`
	// Step name to distinguish multiple compile steps in a build.
	// This is not used when ninja_log is uploaded from chromium developers.
	StepName string `protobuf:"bytes,2,opt,name=step_name,json=stepName,proto3" json:"step_name,omitempty"`
	// ninja’s -j value
	Jobs int64        `protobuf:"varint,3,opt,name=jobs,proto3" json:"jobs,omitempty"`
	Os   NinjaTask_OS `protobuf:"varint,11,opt,name=os,proto3,enum=ninjaproto.NinjaTask_OS" json:"os,omitempty"`
	// The number of cpu cores.
	CpuCore      int32                 `protobuf:"varint,12,opt,name=cpu_core,json=cpuCore,proto3" json:"cpu_core,omitempty"`
	BuildConfigs []*NinjaTask_KeyValue `protobuf:"bytes,13,rep,name=build_configs,json=buildConfigs,proto3" json:"build_configs,omitempty"`
	LogEntry     *NinjaTask_LogEntry   `protobuf:"bytes,4,opt,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	// Weighted build time. This lets us know less parallelized slow build tasks.
	// More details in https://chromium.googlesource.com/chromium/tools/depot_tools/+/5888d6f676722fdac3f65e673c0232667309296c/post_build_ninja_summary.py#52
	// We don't use "google.protobuf.Duration" here in order to make it a bit easier to write a query.
	WeightedDurationSec float64 `protobuf:"fixed64,6,opt,name=weighted_duration_sec,json=weightedDurationSec,proto3" json:"weighted_duration_sec,omitempty"`
	// created_at indicates the timestamp when the NinjaTask is uploaded to BigQuery.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NinjaTask) Reset()         { *m = NinjaTask{} }
func (m *NinjaTask) String() string { return proto.CompactTextString(m) }
func (*NinjaTask) ProtoMessage()    {}
func (*NinjaTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_11621a7b7e2b33ea, []int{0}
}

func (m *NinjaTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NinjaTask.Unmarshal(m, b)
}
func (m *NinjaTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NinjaTask.Marshal(b, m, deterministic)
}
func (m *NinjaTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NinjaTask.Merge(m, src)
}
func (m *NinjaTask) XXX_Size() int {
	return xxx_messageInfo_NinjaTask.Size(m)
}
func (m *NinjaTask) XXX_DiscardUnknown() {
	xxx_messageInfo_NinjaTask.DiscardUnknown(m)
}

var xxx_messageInfo_NinjaTask proto.InternalMessageInfo

func (m *NinjaTask) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *NinjaTask) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *NinjaTask) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *NinjaTask) GetJobs() int64 {
	if m != nil {
		return m.Jobs
	}
	return 0
}

func (m *NinjaTask) GetOs() NinjaTask_OS {
	if m != nil {
		return m.Os
	}
	return NinjaTask_UNKNOWN
}

func (m *NinjaTask) GetCpuCore() int32 {
	if m != nil {
		return m.CpuCore
	}
	return 0
}

func (m *NinjaTask) GetBuildConfigs() []*NinjaTask_KeyValue {
	if m != nil {
		return m.BuildConfigs
	}
	return nil
}

func (m *NinjaTask) GetLogEntry() *NinjaTask_LogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func (m *NinjaTask) GetWeightedDurationSec() float64 {
	if m != nil {
		return m.WeightedDurationSec
	}
	return 0
}

func (m *NinjaTask) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// Content of whitelisted args.gn.
type NinjaTask_KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NinjaTask_KeyValue) Reset()         { *m = NinjaTask_KeyValue{} }
func (m *NinjaTask_KeyValue) String() string { return proto.CompactTextString(m) }
func (*NinjaTask_KeyValue) ProtoMessage()    {}
func (*NinjaTask_KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_11621a7b7e2b33ea, []int{0, 0}
}

func (m *NinjaTask_KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NinjaTask_KeyValue.Unmarshal(m, b)
}
func (m *NinjaTask_KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NinjaTask_KeyValue.Marshal(b, m, deterministic)
}
func (m *NinjaTask_KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NinjaTask_KeyValue.Merge(m, src)
}
func (m *NinjaTask_KeyValue) XXX_Size() int {
	return xxx_messageInfo_NinjaTask_KeyValue.Size(m)
}
func (m *NinjaTask_KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NinjaTask_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_NinjaTask_KeyValue proto.InternalMessageInfo

func (m *NinjaTask_KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NinjaTask_KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Ninja log entry.
// https://github.com/ninja-build/ninja/blob/265a6eaf399778c746c7d2c02b8742f3c0ff07e9/src/build_log.h#L54
type NinjaTask_LogEntry struct {
	// Output filenames of task grouped by command_hash.
	Outputs []string `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Hash of running command.
	CommandHash string `protobuf:"bytes,2,opt,name=command_hash,json=commandHash,proto3" json:"command_hash,omitempty"`
	// Duration between the time starting task and the time when ninja started.
	// We don't use "google.protobuf.Duration" here in order to make it a bit easier to write a query.
	StartDurationSec float64 `protobuf:"fixed64,5,opt,name=start_duration_sec,json=startDurationSec,proto3" json:"start_duration_sec,omitempty"`
	// Duration between the time ending task and the time when ninja started.
	// We don't use "google.protobuf.Duration" here in order to make it a bit easier to write a query.
	EndDurationSec       float64  `protobuf:"fixed64,6,opt,name=end_duration_sec,json=endDurationSec,proto3" json:"end_duration_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NinjaTask_LogEntry) Reset()         { *m = NinjaTask_LogEntry{} }
func (m *NinjaTask_LogEntry) String() string { return proto.CompactTextString(m) }
func (*NinjaTask_LogEntry) ProtoMessage()    {}
func (*NinjaTask_LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_11621a7b7e2b33ea, []int{0, 1}
}

func (m *NinjaTask_LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NinjaTask_LogEntry.Unmarshal(m, b)
}
func (m *NinjaTask_LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NinjaTask_LogEntry.Marshal(b, m, deterministic)
}
func (m *NinjaTask_LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NinjaTask_LogEntry.Merge(m, src)
}
func (m *NinjaTask_LogEntry) XXX_Size() int {
	return xxx_messageInfo_NinjaTask_LogEntry.Size(m)
}
func (m *NinjaTask_LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_NinjaTask_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_NinjaTask_LogEntry proto.InternalMessageInfo

func (m *NinjaTask_LogEntry) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *NinjaTask_LogEntry) GetCommandHash() string {
	if m != nil {
		return m.CommandHash
	}
	return ""
}

func (m *NinjaTask_LogEntry) GetStartDurationSec() float64 {
	if m != nil {
		return m.StartDurationSec
	}
	return 0
}

func (m *NinjaTask_LogEntry) GetEndDurationSec() float64 {
	if m != nil {
		return m.EndDurationSec
	}
	return 0
}

func init() {
	proto.RegisterEnum("ninjaproto.NinjaTask_OS", NinjaTask_OS_name, NinjaTask_OS_value)
	proto.RegisterType((*NinjaTask)(nil), "ninjaproto.NinjaTask")
	proto.RegisterType((*NinjaTask_KeyValue)(nil), "ninjaproto.NinjaTask.KeyValue")
	proto.RegisterType((*NinjaTask_LogEntry)(nil), "ninjaproto.NinjaTask.LogEntry")
}

func init() {
	proto.RegisterFile("ninjalog.proto", fileDescriptor_11621a7b7e2b33ea)
}

var fileDescriptor_11621a7b7e2b33ea = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x6f, 0x94, 0x40,
	0x14, 0x95, 0x05, 0xca, 0x70, 0xb7, 0x6d, 0x26, 0xa3, 0x26, 0xe3, 0x9a, 0x28, 0xf6, 0x89, 0x07,
	0x43, 0x93, 0xf5, 0xc9, 0xf8, 0xd4, 0xac, 0x26, 0x76, 0x5b, 0x69, 0x32, 0x6d, 0xad, 0x6f, 0x64,
	0x16, 0xa6, 0x2c, 0x2d, 0x30, 0x84, 0x19, 0x34, 0xfb, 0x9b, 0xfc, 0x53, 0xfe, 0x14, 0xc3, 0x00,
	0x5a, 0x93, 0xfa, 0x76, 0xcf, 0xfd, 0x9a, 0x73, 0xee, 0x19, 0x38, 0xac, 0x8b, 0xfa, 0x8e, 0x97,
	0x32, 0x8f, 0x9a, 0x56, 0x6a, 0x49, 0xc0, 0x60, 0x13, 0x2f, 0x5e, 0xe7, 0x52, 0xe6, 0xa5, 0x38,
	0x36, 0x68, 0xd3, 0xdd, 0x1e, 0xeb, 0xa2, 0x12, 0x4a, 0xf3, 0xaa, 0x19, 0x9a, 0x8f, 0x7e, 0xb9,
	0xe0, 0xc7, 0x7d, 0xff, 0x15, 0x57, 0xf7, 0xe4, 0x05, 0xa0, 0x4d, 0x57, 0x94, 0x59, 0x52, 0x64,
	0xd4, 0x0a, 0xac, 0xd0, 0x66, 0x9e, 0xc1, 0xa7, 0x19, 0xa1, 0xe0, 0x69, 0xde, 0xe6, 0x42, 0x2b,
	0x8a, 0x02, 0x3b, 0xf4, 0xd9, 0x04, 0xc9, 0x4b, 0xf0, 0x95, 0x16, 0x4d, 0x52, 0xf3, 0x4a, 0xd0,
	0x59, 0x60, 0x85, 0x3e, 0x43, 0x7d, 0x22, 0xe6, 0x95, 0x20, 0x04, 0x9c, 0x3b, 0xb9, 0x51, 0xd4,
	0x36, 0xdb, 0x4c, 0x4c, 0x42, 0x98, 0x49, 0x45, 0xe7, 0x81, 0x15, 0x1e, 0x2e, 0x69, 0xf4, 0x97,
	0x6d, 0xf4, 0x87, 0x48, 0x74, 0x71, 0xc9, 0x66, 0x52, 0xf5, 0x7c, 0xd2, 0xa6, 0x4b, 0x52, 0xd9,
	0x0a, 0xba, 0x1f, 0x58, 0xa1, 0xcb, 0xbc, 0xb4, 0xe9, 0x56, 0xb2, 0x15, 0x64, 0x05, 0x07, 0x03,
	0xd5, 0x54, 0xd6, 0xb7, 0x45, 0xae, 0xe8, 0x41, 0x60, 0x87, 0xf3, 0xe5, 0xab, 0xc7, 0xf7, 0x9d,
	0x89, 0xdd, 0x57, 0x5e, 0x76, 0x82, 0xed, 0x9b, 0xa1, 0xd5, 0x30, 0x43, 0x3e, 0x80, 0x5f, 0xca,
	0x3c, 0x11, 0xb5, 0x6e, 0x77, 0xd4, 0x09, 0xac, 0xff, 0x2f, 0x38, 0x97, 0xf9, 0xa7, 0xbe, 0x8b,
	0xa1, 0x72, 0x8c, 0xc8, 0x12, 0x9e, 0xff, 0x10, 0x45, 0xbe, 0xd5, 0x22, 0x4b, 0xb2, 0xae, 0xe5,
	0xba, 0x90, 0x75, 0xa2, 0x44, 0x4a, 0xf7, 0x02, 0x2b, 0xb4, 0xd8, 0xd3, 0xa9, 0xf8, 0x71, 0xac,
	0x5d, 0x8a, 0x94, 0xbc, 0x07, 0x48, 0x5b, 0xc1, 0xfb, 0x11, 0xae, 0xa9, 0x67, 0x5e, 0x5c, 0x44,
	0x83, 0x49, 0xd1, 0x64, 0x52, 0x74, 0x35, 0x99, 0xc4, 0xfc, 0xb1, 0xfb, 0x44, 0x2f, 0x96, 0x80,
	0x26, 0x15, 0x04, 0x83, 0x7d, 0x2f, 0x76, 0xc6, 0x22, 0x9f, 0xf5, 0x21, 0x79, 0x06, 0xee, 0xf7,
	0xbe, 0x34, 0x1a, 0x30, 0x80, 0xc5, 0x4f, 0x0b, 0xd0, 0xc4, 0xbc, 0x77, 0x50, 0x76, 0xba, 0xe9,
	0xb4, 0xa2, 0xd6, 0xe0, 0xe0, 0x08, 0xc9, 0x1b, 0xd8, 0x4f, 0x65, 0x55, 0xf1, 0x3a, 0x4b, 0xb6,
	0x5c, 0x6d, 0xc7, 0x1d, 0xf3, 0x31, 0xf7, 0x99, 0xab, 0x2d, 0x79, 0x0b, 0x44, 0x69, 0xde, 0xea,
	0x7f, 0x95, 0xba, 0x46, 0x29, 0x36, 0x95, 0x87, 0x32, 0x43, 0xc0, 0xa2, 0x7e, 0xf4, 0x2a, 0x87,
	0xa2, 0x7e, 0x78, 0x90, 0xb5, 0x83, 0x6c, 0xec, 0xac, 0x1d, 0xe4, 0x60, 0xf7, 0x28, 0x82, 0xd9,
	0xc5, 0x25, 0x99, 0x83, 0x77, 0x1d, 0x9f, 0xc5, 0x17, 0x37, 0x31, 0x7e, 0x42, 0x3c, 0xb0, 0x6f,
	0x4e, 0x63, 0x6c, 0x11, 0x1f, 0xdc, 0xf3, 0xd3, 0xf8, 0xfa, 0x1b, 0x9e, 0xf5, 0xb9, 0x2f, 0x27,
	0x2b, 0x6c, 0xaf, 0x1d, 0xe4, 0x63, 0x58, 0x3b, 0x08, 0xf0, 0x7c, 0xed, 0x20, 0x17, 0xef, 0x31,
	0x7f, 0xf8, 0x12, 0x59, 0xd1, 0x32, 0xb4, 0x95, 0x4a, 0xf7, 0x5f, 0x72, 0xb3, 0x67, 0xee, 0xfa,
	0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x81, 0x37, 0xc6, 0x28, 0x03, 0x00, 0x00,
}
