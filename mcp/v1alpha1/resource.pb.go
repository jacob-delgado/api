// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: mcp/v1alpha1/resource.proto

// This package defines the common, core types used by the Mesh Configuration Protocol.

package v1alpha1

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

// Resource as transferred via the Mesh Configuration Protocol. Each
// resource is made up of common metadata, and a type-specific resource payload.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common metadata describing the resource.
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The primary payload for the resource.
	Body *any1.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcp_v1alpha1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_mcp_v1alpha1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_mcp_v1alpha1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Resource) GetBody() *any1.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_mcp_v1alpha1_resource_proto protoreflect.FileDescriptor

var file_mcp_v1alpha1_resource_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x63,
	0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x6d, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x63, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcp_v1alpha1_resource_proto_rawDescOnce sync.Once
	file_mcp_v1alpha1_resource_proto_rawDescData = file_mcp_v1alpha1_resource_proto_rawDesc
)

func file_mcp_v1alpha1_resource_proto_rawDescGZIP() []byte {
	file_mcp_v1alpha1_resource_proto_rawDescOnce.Do(func() {
		file_mcp_v1alpha1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcp_v1alpha1_resource_proto_rawDescData)
	})
	return file_mcp_v1alpha1_resource_proto_rawDescData
}

var file_mcp_v1alpha1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcp_v1alpha1_resource_proto_goTypes = []interface{}{
	(*Resource)(nil), // 0: istio.mcp.v1alpha1.Resource
	(*Metadata)(nil), // 1: istio.mcp.v1alpha1.Metadata
	(*any1.Any)(nil), // 2: google.protobuf.Any
}
var file_mcp_v1alpha1_resource_proto_depIdxs = []int32{
	1, // 0: istio.mcp.v1alpha1.Resource.metadata:type_name -> istio.mcp.v1alpha1.Metadata
	2, // 1: istio.mcp.v1alpha1.Resource.body:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcp_v1alpha1_resource_proto_init() }
func file_mcp_v1alpha1_resource_proto_init() {
	if File_mcp_v1alpha1_resource_proto != nil {
		return
	}
	file_mcp_v1alpha1_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcp_v1alpha1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_mcp_v1alpha1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcp_v1alpha1_resource_proto_goTypes,
		DependencyIndexes: file_mcp_v1alpha1_resource_proto_depIdxs,
		MessageInfos:      file_mcp_v1alpha1_resource_proto_msgTypes,
	}.Build()
	File_mcp_v1alpha1_resource_proto = out.File
	file_mcp_v1alpha1_resource_proto_rawDesc = nil
	file_mcp_v1alpha1_resource_proto_goTypes = nil
	file_mcp_v1alpha1_resource_proto_depIdxs = nil
}