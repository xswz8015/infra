// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/unifiedfleet/api/v1/models/chromeos/lab/chameleon.proto

package ufspb

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

type ChameleonType int32

const (
	ChameleonType_CHAMELEON_TYPE_INVALID      ChameleonType = 0
	ChameleonType_CHAMELEON_TYPE_BT_HID       ChameleonType = 1
	ChameleonType_CHAMELEON_TYPE_DP           ChameleonType = 2
	ChameleonType_CHAMELEON_TYPE_DP_HDMI      ChameleonType = 3
	ChameleonType_CHAMELEON_TYPE_VGA          ChameleonType = 4
	ChameleonType_CHAMELEON_TYPE_HDMI         ChameleonType = 5
	ChameleonType_CHAMELEON_TYPE_BT_BLE_HID   ChameleonType = 6
	ChameleonType_CHAMELEON_TYPE_BT_A2DP_SINK ChameleonType = 7
	ChameleonType_CHAMELEON_TYPE_BT_PEER      ChameleonType = 8
)

// Enum value maps for ChameleonType.
var (
	ChameleonType_name = map[int32]string{
		0: "CHAMELEON_TYPE_INVALID",
		1: "CHAMELEON_TYPE_BT_HID",
		2: "CHAMELEON_TYPE_DP",
		3: "CHAMELEON_TYPE_DP_HDMI",
		4: "CHAMELEON_TYPE_VGA",
		5: "CHAMELEON_TYPE_HDMI",
		6: "CHAMELEON_TYPE_BT_BLE_HID",
		7: "CHAMELEON_TYPE_BT_A2DP_SINK",
		8: "CHAMELEON_TYPE_BT_PEER",
	}
	ChameleonType_value = map[string]int32{
		"CHAMELEON_TYPE_INVALID":      0,
		"CHAMELEON_TYPE_BT_HID":       1,
		"CHAMELEON_TYPE_DP":           2,
		"CHAMELEON_TYPE_DP_HDMI":      3,
		"CHAMELEON_TYPE_VGA":          4,
		"CHAMELEON_TYPE_HDMI":         5,
		"CHAMELEON_TYPE_BT_BLE_HID":   6,
		"CHAMELEON_TYPE_BT_A2DP_SINK": 7,
		"CHAMELEON_TYPE_BT_PEER":      8,
	}
)

func (x ChameleonType) Enum() *ChameleonType {
	p := new(ChameleonType)
	*p = x
	return p
}

func (x ChameleonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChameleonType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_enumTypes[0].Descriptor()
}

func (ChameleonType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_enumTypes[0]
}

func (x ChameleonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChameleonType.Descriptor instead.
func (ChameleonType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescGZIP(), []int{0}
}

// NEXT TAG: 4
type Chameleon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChameleonPeripherals []ChameleonType `protobuf:"varint,3,rep,packed,name=chameleon_peripherals,json=chameleonPeripherals,proto3,enum=unifiedfleet.api.v1.models.chromeos.lab.ChameleonType" json:"chameleon_peripherals,omitempty"`
	// Indicate if there's audio_board in the chameleon.
	AudioBoard bool `protobuf:"varint,2,opt,name=audio_board,json=audioBoard,proto3" json:"audio_board,omitempty"`
}

func (x *Chameleon) Reset() {
	*x = Chameleon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chameleon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chameleon) ProtoMessage() {}

func (x *Chameleon) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chameleon.ProtoReflect.Descriptor instead.
func (*Chameleon) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescGZIP(), []int{0}
}

func (x *Chameleon) GetChameleonPeripherals() []ChameleonType {
	if x != nil {
		return x.ChameleonPeripherals
	}
	return nil
}

func (x *Chameleon) GetAudioBoard() bool {
	if x != nil {
		return x.AudioBoard
	}
	return false
}

var File_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x2f,
	0x63, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6d, 0x65, 0x6c,
	0x65, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62, 0x2e,
	0x43, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x63,
	0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x61, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x2a, 0x86, 0x02, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x48, 0x41, 0x4d, 0x45, 0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x48, 0x41, 0x4d,
	0x45, 0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x54, 0x5f, 0x48, 0x49,
	0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x4d, 0x45, 0x4c, 0x45, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x50, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48,
	0x41, 0x4d, 0x45, 0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x50, 0x5f,
	0x48, 0x44, 0x4d, 0x49, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x41, 0x4d, 0x45, 0x4c,
	0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x47, 0x41, 0x10, 0x04, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x48, 0x41, 0x4d, 0x45, 0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x48, 0x44, 0x4d, 0x49, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x48, 0x41, 0x4d, 0x45,
	0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x54, 0x5f, 0x42, 0x4c, 0x45,
	0x5f, 0x48, 0x49, 0x44, 0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x48, 0x41, 0x4d, 0x45, 0x4c,
	0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x54, 0x5f, 0x41, 0x32, 0x44, 0x50,
	0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x41, 0x4d, 0x45,
	0x4c, 0x45, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x54, 0x5f, 0x50, 0x45, 0x45,
	0x52, 0x10, 0x08, 0x42, 0x35, 0x5a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73,
	0x2f, 0x6c, 0x61, 0x62, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_goTypes = []interface{}{
	(ChameleonType)(0), // 0: unifiedfleet.api.v1.models.chromeos.lab.ChameleonType
	(*Chameleon)(nil),  // 1: unifiedfleet.api.v1.models.chromeos.lab.Chameleon
}
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.chromeos.lab.Chameleon.chameleon_peripherals:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ChameleonType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_init() }
func file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chameleon); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto = out.File
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_chameleon_proto_depIdxs = nil
}
