// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/unifiedfleet/api/v1/models/rack_lse.proto

package ufspb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RackLSE is the Rack Lab Setup Environment.
// It refers to the entity in the lab which has Rack(s) associated with it.
// It also has other components associated with it like switches, kvms, rpms.
type RackLSE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique random generated string
	// The format will be rackLSEs/XXX
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The prototype that this rack LSE should follow. System will use this
	// prototype to detect if the LSE is completed or valid.
	RackLsePrototype string `protobuf:"bytes,2,opt,name=rack_lse_prototype,json=rackLsePrototype,proto3" json:"rack_lse_prototype,omitempty"`
	// Types that are assignable to Lse:
	//	*RackLSE_ChromeBrowserRackLse
	//	*RackLSE_ChromeosRackLse
	Lse isRackLSE_Lse `protobuf_oneof:"lse"`
	// The racks that this LSE is linked to. No rack is linked if it's NULL.
	Racks []string `protobuf:"bytes,5,rep,name=racks,proto3" json:"racks,omitempty"`
	// Record the last update timestamp of this RackLSE (In UTC timezone)
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *RackLSE) Reset() {
	*x = RackLSE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RackLSE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RackLSE) ProtoMessage() {}

func (x *RackLSE) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RackLSE.ProtoReflect.Descriptor instead.
func (*RackLSE) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescGZIP(), []int{0}
}

func (x *RackLSE) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RackLSE) GetRackLsePrototype() string {
	if x != nil {
		return x.RackLsePrototype
	}
	return ""
}

func (m *RackLSE) GetLse() isRackLSE_Lse {
	if m != nil {
		return m.Lse
	}
	return nil
}

func (x *RackLSE) GetChromeBrowserRackLse() *ChromeBrowserRackLSE {
	if x, ok := x.GetLse().(*RackLSE_ChromeBrowserRackLse); ok {
		return x.ChromeBrowserRackLse
	}
	return nil
}

func (x *RackLSE) GetChromeosRackLse() *ChromeOSRackLSE {
	if x, ok := x.GetLse().(*RackLSE_ChromeosRackLse); ok {
		return x.ChromeosRackLse
	}
	return nil
}

func (x *RackLSE) GetRacks() []string {
	if x != nil {
		return x.Racks
	}
	return nil
}

func (x *RackLSE) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type isRackLSE_Lse interface {
	isRackLSE_Lse()
}

type RackLSE_ChromeBrowserRackLse struct {
	ChromeBrowserRackLse *ChromeBrowserRackLSE `protobuf:"bytes,3,opt,name=chrome_browser_rack_lse,json=chromeBrowserRackLse,proto3,oneof"`
}

type RackLSE_ChromeosRackLse struct {
	ChromeosRackLse *ChromeOSRackLSE `protobuf:"bytes,4,opt,name=chromeos_rack_lse,json=chromeosRackLse,proto3,oneof"`
}

func (*RackLSE_ChromeBrowserRackLse) isRackLSE_Lse() {}

func (*RackLSE_ChromeosRackLse) isRackLSE_Lse() {}

// ChromeBrowserRackLSE refers to the entity which has
// kvms and rpms connected to it in Chrome Browser lab
type ChromeBrowserRackLSE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KVMs in the rack, they're the attached kvms' names, which are the same as their hostnames
	//
	// Deprecated: Do not use.
	Kvms []string `protobuf:"bytes,1,rep,name=kvms,proto3" json:"kvms,omitempty"`
	// Switches in the rack.
	//
	// Deprecated: Do not use.
	Switches []string `protobuf:"bytes,2,rep,name=switches,proto3" json:"switches,omitempty"`
}

func (x *ChromeBrowserRackLSE) Reset() {
	*x = ChromeBrowserRackLSE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChromeBrowserRackLSE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChromeBrowserRackLSE) ProtoMessage() {}

func (x *ChromeBrowserRackLSE) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChromeBrowserRackLSE.ProtoReflect.Descriptor instead.
func (*ChromeBrowserRackLSE) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *ChromeBrowserRackLSE) GetKvms() []string {
	if x != nil {
		return x.Kvms
	}
	return nil
}

// Deprecated: Do not use.
func (x *ChromeBrowserRackLSE) GetSwitches() []string {
	if x != nil {
		return x.Switches
	}
	return nil
}

// ChromeOSRackLSE refers to the entity which has
// switches and rpms connected to it in Chrome OS lab
type ChromeOSRackLSE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RPMs in the rack
	Rpms []string `protobuf:"bytes,1,rep,name=rpms,proto3" json:"rpms,omitempty"`
	// KVMs in the rack
	Kvms     []string `protobuf:"bytes,2,rep,name=kvms,proto3" json:"kvms,omitempty"`
	Switches []string `protobuf:"bytes,3,rep,name=switches,proto3" json:"switches,omitempty"`
}

func (x *ChromeOSRackLSE) Reset() {
	*x = ChromeOSRackLSE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChromeOSRackLSE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChromeOSRackLSE) ProtoMessage() {}

func (x *ChromeOSRackLSE) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChromeOSRackLSE.ProtoReflect.Descriptor instead.
func (*ChromeOSRackLSE) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescGZIP(), []int{2}
}

func (x *ChromeOSRackLSE) GetRpms() []string {
	if x != nil {
		return x.Rpms
	}
	return nil
}

func (x *ChromeOSRackLSE) GetKvms() []string {
	if x != nil {
		return x.Kvms
	}
	return nil
}

func (x *ChromeOSRackLSE) GetSwitches() []string {
	if x != nil {
		return x.Switches
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_rack_lse_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x04, 0x0a, 0x07, 0x52,
	0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x12, 0x72, 0x61,
	0x63, 0x6b, 0x5f, 0x6c, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0xfa, 0x41, 0x33, 0x0a, 0x31, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61,
	0x63, 0x6b, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0xe0, 0x41,
	0x02, 0x52, 0x10, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x69, 0x0a, 0x17, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x52,
	0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x48, 0x00, 0x52, 0x14, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x11, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x72, 0x61, 0x63, 0x6b, 0x5f,
	0x6c, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x4f, 0x53, 0x52,
	0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2a, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x52, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x41, 0xea,
	0x41, 0x3e, 0x0a, 0x28, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x12, 0x12, 0x72, 0x61,
	0x63, 0x6b, 0x4c, 0x53, 0x45, 0x73, 0x2f, 0x7b, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x7d,
	0x42, 0x05, 0x0a, 0x03, 0x6c, 0x73, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45,
	0x12, 0x3f, 0x0a, 0x04, 0x6b, 0x76, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2b,
	0x18, 0x01, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x56, 0x4d, 0x52, 0x04, 0x6b, 0x76, 0x6d,
	0x73, 0x12, 0x4a, 0x0a, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x2e, 0x18, 0x01, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x52, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0xd9, 0x01,
	0x0a, 0x0f, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x4f, 0x53, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x53,
	0x45, 0x12, 0x3d, 0x0a, 0x04, 0x72, 0x70, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x29, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70,
	0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x50, 0x4d, 0x52, 0x04, 0x72, 0x70, 0x6d, 0x73,
	0x12, 0x3d, 0x0a, 0x04, 0x6b, 0x76, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x29,
	0xfa, 0x41, 0x26, 0x0a, 0x24, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x56, 0x4d, 0x52, 0x04, 0x6b, 0x76, 0x6d, 0x73, 0x12,
	0x48, 0x0a, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x2c, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52,
	0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_unifiedfleet_api_v1_models_rack_lse_proto_goTypes = []interface{}{
	(*RackLSE)(nil),              // 0: unifiedfleet.api.v1.models.RackLSE
	(*ChromeBrowserRackLSE)(nil), // 1: unifiedfleet.api.v1.models.ChromeBrowserRackLSE
	(*ChromeOSRackLSE)(nil),      // 2: unifiedfleet.api.v1.models.ChromeOSRackLSE
	(*timestamp.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_infra_unifiedfleet_api_v1_models_rack_lse_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.RackLSE.chrome_browser_rack_lse:type_name -> unifiedfleet.api.v1.models.ChromeBrowserRackLSE
	2, // 1: unifiedfleet.api.v1.models.RackLSE.chromeos_rack_lse:type_name -> unifiedfleet.api.v1.models.ChromeOSRackLSE
	3, // 2: unifiedfleet.api.v1.models.RackLSE.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_rack_lse_proto_init() }
func file_infra_unifiedfleet_api_v1_models_rack_lse_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_rack_lse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RackLSE); i {
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
		file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChromeBrowserRackLSE); i {
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
		file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChromeOSRackLSE); i {
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
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RackLSE_ChromeBrowserRackLse)(nil),
		(*RackLSE_ChromeosRackLse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_rack_lse_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_rack_lse_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_rack_lse_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_rack_lse_proto = out.File
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_rack_lse_proto_depIdxs = nil
}
