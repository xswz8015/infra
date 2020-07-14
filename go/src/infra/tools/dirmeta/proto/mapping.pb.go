// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/tools/dirmeta/proto/mapping.proto

package dirmetapb

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

// A known form of the Mapping.
type MappingForm int32

const (
	// The mapping represents the metadata files as is.
	// In particular, a Dirs entry exists iff the directory defines metadata.
	// The metadata of the entry does not include inherited metadata.
	MappingForm_ORIGINAL MappingForm = 0
	// Like ORIGINAL, but each entry includes inherited metadata.
	MappingForm_COMPUTED MappingForm = 1
	// Like COMPUTED, but a Dirs entry exists even if the directory does not
	// define any metadata.
	MappingForm_FULL MappingForm = 2
	// The mapping contains no redundant information.
	// For example, it is redundant to specify the same value for the same
	// metadata attribute in both "a" and "a/b".
	// It is also redundant to have an Dirs entry with empty metadata.
	// This is the most compact form without a dataloss.
	MappingForm_REDUCED MappingForm = 3
)

// Enum value maps for MappingForm.
var (
	MappingForm_name = map[int32]string{
		0: "ORIGINAL",
		1: "COMPUTED",
		2: "FULL",
		3: "REDUCED",
	}
	MappingForm_value = map[string]int32{
		"ORIGINAL": 0,
		"COMPUTED": 1,
		"FULL":     2,
		"REDUCED":  3,
	}
)

func (x MappingForm) Enum() *MappingForm {
	p := new(MappingForm)
	*p = x
	return p
}

func (x MappingForm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MappingForm) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_tools_dirmeta_proto_mapping_proto_enumTypes[0].Descriptor()
}

func (MappingForm) Type() protoreflect.EnumType {
	return &file_infra_tools_dirmeta_proto_mapping_proto_enumTypes[0]
}

func (x MappingForm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MappingForm.Descriptor instead.
func (MappingForm) EnumDescriptor() ([]byte, []int) {
	return file_infra_tools_dirmeta_proto_mapping_proto_rawDescGZIP(), []int{0}
}

// Maps from a directory to its metadata.
type Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dirs maps from a directory to its metadata.
	//
	// The key is directory name, relative to the root.
	// The key must use forward slash as directory separator.
	// The key must be clean: https://pkg.go.dev/path?tab=doc#Clean
	// Special key "." represents the root directory.
	//
	// The root must be known from the context where Mapping is used and is not
	// part of the this message.
	Dirs map[string]*Metadata `protobuf:"bytes,1,rep,name=dirs,proto3" json:"dirs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Mapping) Reset() {
	*x = Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tools_dirmeta_proto_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mapping) ProtoMessage() {}

func (x *Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tools_dirmeta_proto_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mapping.ProtoReflect.Descriptor instead.
func (*Mapping) Descriptor() ([]byte, []int) {
	return file_infra_tools_dirmeta_proto_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *Mapping) GetDirs() map[string]*Metadata {
	if x != nil {
		return x.Dirs
	}
	return nil
}

var File_infra_tools_dirmeta_proto_mapping_proto protoreflect.FileDescriptor

var file_infra_tools_dirmeta_proto_mapping_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x69,
	0x72, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x65, 0x2e, 0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2c,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x69, 0x72, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x07, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x69, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e,
	0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x69, 0x72, 0x73, 0x1a, 0x56, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x64, 0x69, 0x72, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x40, 0x0a, 0x0b,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0c, 0x0a, 0x08, 0x4f,
	0x52, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d,
	0x50, 0x55, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x44, 0x10, 0x03, 0x42, 0x25,
	0x5a, 0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x69,
	0x72, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x64, 0x69, 0x72, 0x6d,
	0x65, 0x74, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tools_dirmeta_proto_mapping_proto_rawDescOnce sync.Once
	file_infra_tools_dirmeta_proto_mapping_proto_rawDescData = file_infra_tools_dirmeta_proto_mapping_proto_rawDesc
)

func file_infra_tools_dirmeta_proto_mapping_proto_rawDescGZIP() []byte {
	file_infra_tools_dirmeta_proto_mapping_proto_rawDescOnce.Do(func() {
		file_infra_tools_dirmeta_proto_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tools_dirmeta_proto_mapping_proto_rawDescData)
	})
	return file_infra_tools_dirmeta_proto_mapping_proto_rawDescData
}

var file_infra_tools_dirmeta_proto_mapping_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_tools_dirmeta_proto_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tools_dirmeta_proto_mapping_proto_goTypes = []interface{}{
	(MappingForm)(0), // 0: chrome.dir_metadata.MappingForm
	(*Mapping)(nil),  // 1: chrome.dir_metadata.Mapping
	nil,              // 2: chrome.dir_metadata.Mapping.DirsEntry
	(*Metadata)(nil), // 3: chrome.dir_metadata.Metadata
}
var file_infra_tools_dirmeta_proto_mapping_proto_depIdxs = []int32{
	2, // 0: chrome.dir_metadata.Mapping.dirs:type_name -> chrome.dir_metadata.Mapping.DirsEntry
	3, // 1: chrome.dir_metadata.Mapping.DirsEntry.value:type_name -> chrome.dir_metadata.Metadata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_tools_dirmeta_proto_mapping_proto_init() }
func file_infra_tools_dirmeta_proto_mapping_proto_init() {
	if File_infra_tools_dirmeta_proto_mapping_proto != nil {
		return
	}
	file_infra_tools_dirmeta_proto_dir_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_tools_dirmeta_proto_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mapping); i {
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
			RawDescriptor: file_infra_tools_dirmeta_proto_mapping_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tools_dirmeta_proto_mapping_proto_goTypes,
		DependencyIndexes: file_infra_tools_dirmeta_proto_mapping_proto_depIdxs,
		EnumInfos:         file_infra_tools_dirmeta_proto_mapping_proto_enumTypes,
		MessageInfos:      file_infra_tools_dirmeta_proto_mapping_proto_msgTypes,
	}.Build()
	File_infra_tools_dirmeta_proto_mapping_proto = out.File
	file_infra_tools_dirmeta_proto_mapping_proto_rawDesc = nil
	file_infra_tools_dirmeta_proto_mapping_proto_goTypes = nil
	file_infra_tools_dirmeta_proto_mapping_proto_depIdxs = nil
}
