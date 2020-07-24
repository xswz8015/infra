// Copyright 2018 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: common.proto

package inventory

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// NEXT TAG: 4
type Environment int32

const (
	Environment_ENVIRONMENT_INVALID Environment = 0
	Environment_ENVIRONMENT_PROD    Environment = 1
	Environment_ENVIRONMENT_STAGING Environment = 2
	// Deprecated: Do not use.
	Environment_ENVIRONMENT_SKYLAB Environment = 3
)

// Enum value maps for Environment.
var (
	Environment_name = map[int32]string{
		0: "ENVIRONMENT_INVALID",
		1: "ENVIRONMENT_PROD",
		2: "ENVIRONMENT_STAGING",
		3: "ENVIRONMENT_SKYLAB",
	}
	Environment_value = map[string]int32{
		"ENVIRONMENT_INVALID": 0,
		"ENVIRONMENT_PROD":    1,
		"ENVIRONMENT_STAGING": 2,
		"ENVIRONMENT_SKYLAB":  3,
	}
)

func (x Environment) Enum() *Environment {
	p := new(Environment)
	*p = x
	return p
}

func (x Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Environment) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Environment) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Environment(num)
	return nil
}

// Deprecated: Use Environment.Descriptor instead.
func (Environment) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds *int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos   *int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil && x.Seconds != nil {
		return *x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil && x.Nanos != nil {
		return *x.Nanos
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3b, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x2a, 0x71, 0x0a, 0x0b, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x56, 0x49,
	0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x56, 0x49, 0x52,
	0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x12, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x4b, 0x59, 0x4c, 0x41, 0x42, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_proto_goTypes = []interface{}{
	(Environment)(0),  // 0: chrome.chromeos_infra.skylab.proto.inventory.Environment
	(*Timestamp)(nil), // 1: chrome.chromeos_infra.skylab.proto.inventory.Timestamp
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
