// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/gkebackup/v1/common.proto

package gkebackuppb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A list of Kubernetes Namespaces
type Namespaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A list of Kubernetes Namespaces
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *Namespaces) Reset() {
	*x = Namespaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespaces) ProtoMessage() {}

func (x *Namespaces) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespaces.ProtoReflect.Descriptor instead.
func (*Namespaces) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Namespaces) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

// A reference to a namespaced resource in Kubernetes.
type NamespacedName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The Namespace of the Kubernetes resource.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional. The name of the Kubernetes resource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NamespacedName) Reset() {
	*x = NamespacedName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacedName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacedName) ProtoMessage() {}

func (x *NamespacedName) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacedName.ProtoReflect.Descriptor instead.
func (*NamespacedName) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *NamespacedName) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespacedName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A list of namespaced Kubernetes resources.
type NamespacedNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A list of namespaced Kubernetes resources.
	NamespacedNames []*NamespacedName `protobuf:"bytes,1,rep,name=namespaced_names,json=namespacedNames,proto3" json:"namespaced_names,omitempty"`
}

func (x *NamespacedNames) Reset() {
	*x = NamespacedNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacedNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacedNames) ProtoMessage() {}

func (x *NamespacedNames) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacedNames.ProtoReflect.Descriptor instead.
func (*NamespacedNames) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *NamespacedNames) GetNamespacedNames() []*NamespacedName {
	if x != nil {
		return x.NamespacedNames
	}
	return nil
}

// Defined a customer managed encryption key that will be used to encrypt Backup
// artifacts.
type EncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Google Cloud KMS encryption key. Format:
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	GcpKmsEncryptionKey string `protobuf:"bytes,1,opt,name=gcp_kms_encryption_key,json=gcpKmsEncryptionKey,proto3" json:"gcp_kms_encryption_key,omitempty"`
}

func (x *EncryptionKey) Reset() {
	*x = EncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionKey) ProtoMessage() {}

func (x *EncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionKey.ProtoReflect.Descriptor instead.
func (*EncryptionKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptionKey) GetGcpKmsEncryptionKey() string {
	if x != nil {
		return x.GcpKmsEncryptionKey
	}
	return ""
}

var File_google_cloud_gkebackup_v1_common_proto protoreflect.FileDescriptor

var file_google_cloud_gkebackup_v1_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6c, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x6f,
	0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x5e, 0x0a, 0x16, 0x67, 0x63, 0x70, 0x5f, 0x6b, 0x6d, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x01, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x13, 0x67, 0x63, 0x70, 0x4b,
	0x6d, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x42,
	0xc2, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70,
	0x62, 0x3b, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70, 0x62, 0xaa, 0x02, 0x19,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gkebackup_v1_common_proto_rawDescOnce sync.Once
	file_google_cloud_gkebackup_v1_common_proto_rawDescData = file_google_cloud_gkebackup_v1_common_proto_rawDesc
)

func file_google_cloud_gkebackup_v1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_gkebackup_v1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_gkebackup_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gkebackup_v1_common_proto_rawDescData)
	})
	return file_google_cloud_gkebackup_v1_common_proto_rawDescData
}

var file_google_cloud_gkebackup_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_gkebackup_v1_common_proto_goTypes = []interface{}{
	(*Namespaces)(nil),      // 0: google.cloud.gkebackup.v1.Namespaces
	(*NamespacedName)(nil),  // 1: google.cloud.gkebackup.v1.NamespacedName
	(*NamespacedNames)(nil), // 2: google.cloud.gkebackup.v1.NamespacedNames
	(*EncryptionKey)(nil),   // 3: google.cloud.gkebackup.v1.EncryptionKey
}
var file_google_cloud_gkebackup_v1_common_proto_depIdxs = []int32{
	1, // 0: google.cloud.gkebackup.v1.NamespacedNames.namespaced_names:type_name -> google.cloud.gkebackup.v1.NamespacedName
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_gkebackup_v1_common_proto_init() }
func file_google_cloud_gkebackup_v1_common_proto_init() {
	if File_google_cloud_gkebackup_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_gkebackup_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespaces); i {
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
		file_google_cloud_gkebackup_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacedName); i {
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
		file_google_cloud_gkebackup_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacedNames); i {
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
		file_google_cloud_gkebackup_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionKey); i {
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
			RawDescriptor: file_google_cloud_gkebackup_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gkebackup_v1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkebackup_v1_common_proto_depIdxs,
		MessageInfos:      file_google_cloud_gkebackup_v1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_gkebackup_v1_common_proto = out.File
	file_google_cloud_gkebackup_v1_common_proto_rawDesc = nil
	file_google_cloud_gkebackup_v1_common_proto_goTypes = nil
	file_google_cloud_gkebackup_v1_common_proto_depIdxs = nil
}
