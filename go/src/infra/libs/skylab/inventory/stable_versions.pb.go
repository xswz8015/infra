// Copyright 2018 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.1
// source: infra/libs/skylab/inventory/stable_versions.proto

package inventory

import (
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

// NEXT TAG: 5
type StableVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OS image stable versions.
	AndroidOsVersions []*StableVersion `protobuf:"bytes,1,rep,name=android_os_versions,json=androidOsVersions" json:"android_os_versions,omitempty"`
	ChromeOsVersions  []*StableVersion `protobuf:"bytes,2,rep,name=chrome_os_versions,json=chromeOsVersions" json:"chrome_os_versions,omitempty"`
	// Read-write firmware versions. Only relevant for ChromeOS boards.
	RwFirmwareVersions []*StableVersion `protobuf:"bytes,3,rep,name=rw_firmware_versions,json=rwFirmwareVersions" json:"rw_firmware_versions,omitempty"`
	// Used by FAFT testing to install the RO firmware to test. ChromeOS only.
	FaftFirmwareVersions []*StableVersion `protobuf:"bytes,4,rep,name=faft_firmware_versions,json=faftFirmwareVersions" json:"faft_firmware_versions,omitempty"`
}

func (x *StableVersions) Reset() {
	*x = StableVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableVersions) ProtoMessage() {}

func (x *StableVersions) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableVersions.ProtoReflect.Descriptor instead.
func (*StableVersions) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_stable_versions_proto_rawDescGZIP(), []int{0}
}

func (x *StableVersions) GetAndroidOsVersions() []*StableVersion {
	if x != nil {
		return x.AndroidOsVersions
	}
	return nil
}

func (x *StableVersions) GetChromeOsVersions() []*StableVersion {
	if x != nil {
		return x.ChromeOsVersions
	}
	return nil
}

func (x *StableVersions) GetRwFirmwareVersions() []*StableVersion {
	if x != nil {
		return x.RwFirmwareVersions
	}
	return nil
}

func (x *StableVersions) GetFaftFirmwareVersions() []*StableVersion {
	if x != nil {
		return x.FaftFirmwareVersions
	}
	return nil
}

// NEXT TAG: 3
type StableVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board *string `protobuf:"bytes,1,req,name=board" json:"board,omitempty"`
	// Versions are opaque strings for the inventory. Different boards may use the
	// version strings in different ways to obtain the actual images.
	Version *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
}

func (x *StableVersion) Reset() {
	*x = StableVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableVersion) ProtoMessage() {}

func (x *StableVersion) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableVersion.ProtoReflect.Descriptor instead.
func (*StableVersion) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_stable_versions_proto_rawDescGZIP(), []int{1}
}

func (x *StableVersion) GetBoard() string {
	if x != nil && x.Board != nil {
		return *x.Board
	}
	return ""
}

func (x *StableVersion) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

var File_infra_libs_skylab_inventory_stable_versions_proto protoreflect.FileDescriptor

var file_infra_libs_skylab_inventory_stable_versions_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79,
	0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x22, 0xca, 0x03, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6b, 0x0a, 0x13, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x69, 0x0a, 0x12, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x6f, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6d, 0x0a, 0x14,
	0x72, 0x77, 0x5f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x72, 0x77, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x71, 0x0a, 0x16, 0x66,
	0x61, 0x66, 0x74, 0x5f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x66, 0x61, 0x66, 0x74, 0x46, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3f,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x27, 0x5a, 0x25, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b,
	0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x3b, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
}

var (
	file_infra_libs_skylab_inventory_stable_versions_proto_rawDescOnce sync.Once
	file_infra_libs_skylab_inventory_stable_versions_proto_rawDescData = file_infra_libs_skylab_inventory_stable_versions_proto_rawDesc
)

func file_infra_libs_skylab_inventory_stable_versions_proto_rawDescGZIP() []byte {
	file_infra_libs_skylab_inventory_stable_versions_proto_rawDescOnce.Do(func() {
		file_infra_libs_skylab_inventory_stable_versions_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_libs_skylab_inventory_stable_versions_proto_rawDescData)
	})
	return file_infra_libs_skylab_inventory_stable_versions_proto_rawDescData
}

var file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_libs_skylab_inventory_stable_versions_proto_goTypes = []interface{}{
	(*StableVersions)(nil), // 0: chrome.chromeos_infra.skylab.proto.inventory.StableVersions
	(*StableVersion)(nil),  // 1: chrome.chromeos_infra.skylab.proto.inventory.StableVersion
}
var file_infra_libs_skylab_inventory_stable_versions_proto_depIdxs = []int32{
	1, // 0: chrome.chromeos_infra.skylab.proto.inventory.StableVersions.android_os_versions:type_name -> chrome.chromeos_infra.skylab.proto.inventory.StableVersion
	1, // 1: chrome.chromeos_infra.skylab.proto.inventory.StableVersions.chrome_os_versions:type_name -> chrome.chromeos_infra.skylab.proto.inventory.StableVersion
	1, // 2: chrome.chromeos_infra.skylab.proto.inventory.StableVersions.rw_firmware_versions:type_name -> chrome.chromeos_infra.skylab.proto.inventory.StableVersion
	1, // 3: chrome.chromeos_infra.skylab.proto.inventory.StableVersions.faft_firmware_versions:type_name -> chrome.chromeos_infra.skylab.proto.inventory.StableVersion
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infra_libs_skylab_inventory_stable_versions_proto_init() }
func file_infra_libs_skylab_inventory_stable_versions_proto_init() {
	if File_infra_libs_skylab_inventory_stable_versions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableVersions); i {
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
		file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableVersion); i {
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
			RawDescriptor: file_infra_libs_skylab_inventory_stable_versions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_libs_skylab_inventory_stable_versions_proto_goTypes,
		DependencyIndexes: file_infra_libs_skylab_inventory_stable_versions_proto_depIdxs,
		MessageInfos:      file_infra_libs_skylab_inventory_stable_versions_proto_msgTypes,
	}.Build()
	File_infra_libs_skylab_inventory_stable_versions_proto = out.File
	file_infra_libs_skylab_inventory_stable_versions_proto_rawDesc = nil
	file_infra_libs_skylab_inventory_stable_versions_proto_goTypes = nil
	file_infra_libs_skylab_inventory_stable_versions_proto_depIdxs = nil
}
