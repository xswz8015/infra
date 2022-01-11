// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: infra/appengine/chromium_build_stats/ninjaproto/ninjalog.proto

package ninjaproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OS.
type NinjaTask_OS int32

const (
	NinjaTask_UNKNOWN NinjaTask_OS = 0
	NinjaTask_WIN     NinjaTask_OS = 1
	NinjaTask_LINUX   NinjaTask_OS = 2
	NinjaTask_MAC     NinjaTask_OS = 3
)

// Enum value maps for NinjaTask_OS.
var (
	NinjaTask_OS_name = map[int32]string{
		0: "UNKNOWN",
		1: "WIN",
		2: "LINUX",
		3: "MAC",
	}
	NinjaTask_OS_value = map[string]int32{
		"UNKNOWN": 0,
		"WIN":     1,
		"LINUX":   2,
		"MAC":     3,
	}
)

func (x NinjaTask_OS) Enum() *NinjaTask_OS {
	p := new(NinjaTask_OS)
	*p = x
	return p
}

func (x NinjaTask_OS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NinjaTask_OS) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_enumTypes[0].Descriptor()
}

func (NinjaTask_OS) Type() protoreflect.EnumType {
	return &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_enumTypes[0]
}

func (x NinjaTask_OS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NinjaTask_OS.Descriptor instead.
func (NinjaTask_OS) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescGZIP(), []int{0, 0}
}

// NinjaTask message for the log uploaded from chromium developers or buildbot.
// Due to row size limit (1MB) of BigQuery streaming insert, this message
// corresponds to the one task of ninja_log.
// NEXT ID TO USE: 14
type NinjaTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *NinjaTask) Reset() {
	*x = NinjaTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NinjaTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NinjaTask) ProtoMessage() {}

func (x *NinjaTask) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NinjaTask.ProtoReflect.Descriptor instead.
func (*NinjaTask) Descriptor() ([]byte, []int) {
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescGZIP(), []int{0}
}

func (x *NinjaTask) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *NinjaTask) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *NinjaTask) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (x *NinjaTask) GetJobs() int64 {
	if x != nil {
		return x.Jobs
	}
	return 0
}

func (x *NinjaTask) GetOs() NinjaTask_OS {
	if x != nil {
		return x.Os
	}
	return NinjaTask_UNKNOWN
}

func (x *NinjaTask) GetCpuCore() int32 {
	if x != nil {
		return x.CpuCore
	}
	return 0
}

func (x *NinjaTask) GetBuildConfigs() []*NinjaTask_KeyValue {
	if x != nil {
		return x.BuildConfigs
	}
	return nil
}

func (x *NinjaTask) GetLogEntry() *NinjaTask_LogEntry {
	if x != nil {
		return x.LogEntry
	}
	return nil
}

func (x *NinjaTask) GetWeightedDurationSec() float64 {
	if x != nil {
		return x.WeightedDurationSec
	}
	return 0
}

func (x *NinjaTask) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// Content of whitelisted args.gn.
type NinjaTask_KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NinjaTask_KeyValue) Reset() {
	*x = NinjaTask_KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NinjaTask_KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NinjaTask_KeyValue) ProtoMessage() {}

func (x *NinjaTask_KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NinjaTask_KeyValue.ProtoReflect.Descriptor instead.
func (*NinjaTask_KeyValue) Descriptor() ([]byte, []int) {
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NinjaTask_KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NinjaTask_KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Ninja log entry.
// https://github.com/ninja-build/ninja/blob/265a6eaf399778c746c7d2c02b8742f3c0ff07e9/src/build_log.h#L54
type NinjaTask_LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output filenames of task grouped by command_hash.
	Outputs []string `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Hash of running command.
	CommandHash string `protobuf:"bytes,2,opt,name=command_hash,json=commandHash,proto3" json:"command_hash,omitempty"`
	// Duration between the time starting task and the time when ninja started.
	// We don't use "google.protobuf.Duration" here in order to make it a bit easier to write a query.
	StartDurationSec float64 `protobuf:"fixed64,5,opt,name=start_duration_sec,json=startDurationSec,proto3" json:"start_duration_sec,omitempty"`
	// Duration between the time ending task and the time when ninja started.
	// We don't use "google.protobuf.Duration" here in order to make it a bit easier to write a query.
	EndDurationSec float64 `protobuf:"fixed64,6,opt,name=end_duration_sec,json=endDurationSec,proto3" json:"end_duration_sec,omitempty"`
}

func (x *NinjaTask_LogEntry) Reset() {
	*x = NinjaTask_LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NinjaTask_LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NinjaTask_LogEntry) ProtoMessage() {}

func (x *NinjaTask_LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NinjaTask_LogEntry.ProtoReflect.Descriptor instead.
func (*NinjaTask_LogEntry) Descriptor() ([]byte, []int) {
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NinjaTask_LogEntry) GetOutputs() []string {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *NinjaTask_LogEntry) GetCommandHash() string {
	if x != nil {
		return x.CommandHash
	}
	return ""
}

func (x *NinjaTask_LogEntry) GetStartDurationSec() float64 {
	if x != nil {
		return x.StartDurationSec
	}
	return 0
}

func (x *NinjaTask_LogEntry) GetEndDurationSec() float64 {
	if x != nil {
		return x.EndDurationSec
	}
	return 0
}

var File_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto protoreflect.FileDescriptor

var file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x05,
	0x0a, 0x09, 0x4e, 0x69, 0x6e, 0x6a, 0x61, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x12, 0x28, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x69, 0x6e, 0x6a, 0x61,
	0x54, 0x61, 0x73, 0x6b, 0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x69, 0x6e, 0x6a, 0x61,
	0x54, 0x61, 0x73, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x6c,
	0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x69, 0x6e, 0x6a,
	0x61, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xab, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x65,
	0x6e, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x2e, 0x0a, 0x02, 0x4f, 0x53, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x57, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a,
	0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x52, 0x09, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x64, 0x69, 0x72, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x31, 0x5a, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescOnce sync.Once
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescData = file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDesc
)

func file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescGZIP() []byte {
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescOnce.Do(func() {
		file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescData)
	})
	return file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDescData
}

var file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_goTypes = []interface{}{
	(NinjaTask_OS)(0),             // 0: ninjaproto.NinjaTask.OS
	(*NinjaTask)(nil),             // 1: ninjaproto.NinjaTask
	(*NinjaTask_KeyValue)(nil),    // 2: ninjaproto.NinjaTask.KeyValue
	(*NinjaTask_LogEntry)(nil),    // 3: ninjaproto.NinjaTask.LogEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_depIdxs = []int32{
	0, // 0: ninjaproto.NinjaTask.os:type_name -> ninjaproto.NinjaTask.OS
	2, // 1: ninjaproto.NinjaTask.build_configs:type_name -> ninjaproto.NinjaTask.KeyValue
	3, // 2: ninjaproto.NinjaTask.log_entry:type_name -> ninjaproto.NinjaTask.LogEntry
	4, // 3: ninjaproto.NinjaTask.created_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_init() }
func file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_init() {
	if File_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NinjaTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NinjaTask_KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NinjaTask_LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_goTypes,
		DependencyIndexes: file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_depIdxs,
		EnumInfos:         file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_enumTypes,
		MessageInfos:      file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_msgTypes,
	}.Build()
	File_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto = out.File
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_rawDesc = nil
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_goTypes = nil
	file_infra_appengine_chromium_build_stats_ninjaproto_ninjalog_proto_depIdxs = nil
}
