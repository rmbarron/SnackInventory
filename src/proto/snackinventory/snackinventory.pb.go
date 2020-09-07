// Copyright 2020 Robert Barron

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: snackinventory.proto

package snackinventory

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

// A snack is an individual item in our inventory.
// We store a registry of potential snacks, and keep the count of each snack
// currently in inventory.
// Snacks use `barcode` as their unique ID, as multiple different snacks may
// have the same name &/or brand.
type Snack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode string `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Brand   string `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *Snack) Reset() {
	*x = Snack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snack) ProtoMessage() {}

func (x *Snack) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snack.ProtoReflect.Descriptor instead.
func (*Snack) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{0}
}

func (x *Snack) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *Snack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Snack) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

type CreateSnackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snack *Snack `protobuf:"bytes,1,opt,name=snack,proto3" json:"snack,omitempty"`
}

func (x *CreateSnackRequest) Reset() {
	*x = CreateSnackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnackRequest) ProtoMessage() {}

func (x *CreateSnackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnackRequest.ProtoReflect.Descriptor instead.
func (*CreateSnackRequest) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnackRequest) GetSnack() *Snack {
	if x != nil {
		return x.Snack
	}
	return nil
}

// Status / Success is communicated via gRPC response status.
type CreateSnackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSnackResponse) Reset() {
	*x = CreateSnackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnackResponse) ProtoMessage() {}

func (x *CreateSnackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnackResponse.ProtoReflect.Descriptor instead.
func (*CreateSnackResponse) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{2}
}

type ListSnacksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSnacksRequest) Reset() {
	*x = ListSnacksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnacksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnacksRequest) ProtoMessage() {}

func (x *ListSnacksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnacksRequest.ProtoReflect.Descriptor instead.
func (*ListSnacksRequest) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{3}
}

type ListSnacksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snacks []*Snack `protobuf:"bytes,1,rep,name=snacks,proto3" json:"snacks,omitempty"`
}

func (x *ListSnacksResponse) Reset() {
	*x = ListSnacksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnacksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnacksResponse) ProtoMessage() {}

func (x *ListSnacksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnacksResponse.ProtoReflect.Descriptor instead.
func (*ListSnacksResponse) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{4}
}

func (x *ListSnacksResponse) GetSnacks() []*Snack {
	if x != nil {
		return x.Snacks
	}
	return nil
}

// Snacks can only be deleted by barcode, the unique ID for snacks.
type DeleteSnackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode string `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
}

func (x *DeleteSnackRequest) Reset() {
	*x = DeleteSnackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnackRequest) ProtoMessage() {}

func (x *DeleteSnackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnackRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnackRequest) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSnackRequest) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

type DeleteSnackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSnackResponse) Reset() {
	*x = DeleteSnackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snackinventory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnackResponse) ProtoMessage() {}

func (x *DeleteSnackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snackinventory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnackResponse.ProtoReflect.Descriptor instead.
func (*DeleteSnackResponse) Descriptor() ([]byte, []int) {
	return file_snackinventory_proto_rawDescGZIP(), []int{6}
}

var File_snackinventory_proto protoreflect.FileDescriptor

var file_snackinventory_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x4b, 0x0a, 0x05, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x6e, 0x61,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52,
	0x05, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x43, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6e, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52,
	0x06, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9b,
	0x02, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x58, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x63, 0x6b,
	0x12, 0x22, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x6e, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x6e, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x63,
	0x6b, 0x12, 0x22, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6e, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6d, 0x62, 0x61, 0x72,
	0x72, 0x6f, 0x6e, 0x2f, 0x53, 0x6e, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6e, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_snackinventory_proto_rawDescOnce sync.Once
	file_snackinventory_proto_rawDescData = file_snackinventory_proto_rawDesc
)

func file_snackinventory_proto_rawDescGZIP() []byte {
	file_snackinventory_proto_rawDescOnce.Do(func() {
		file_snackinventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_snackinventory_proto_rawDescData)
	})
	return file_snackinventory_proto_rawDescData
}

var file_snackinventory_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_snackinventory_proto_goTypes = []interface{}{
	(*Snack)(nil),               // 0: snackinventory.Snack
	(*CreateSnackRequest)(nil),  // 1: snackinventory.CreateSnackRequest
	(*CreateSnackResponse)(nil), // 2: snackinventory.CreateSnackResponse
	(*ListSnacksRequest)(nil),   // 3: snackinventory.ListSnacksRequest
	(*ListSnacksResponse)(nil),  // 4: snackinventory.ListSnacksResponse
	(*DeleteSnackRequest)(nil),  // 5: snackinventory.DeleteSnackRequest
	(*DeleteSnackResponse)(nil), // 6: snackinventory.DeleteSnackResponse
}
var file_snackinventory_proto_depIdxs = []int32{
	0, // 0: snackinventory.CreateSnackRequest.snack:type_name -> snackinventory.Snack
	0, // 1: snackinventory.ListSnacksResponse.snacks:type_name -> snackinventory.Snack
	1, // 2: snackinventory.SnackInventory.CreateSnack:input_type -> snackinventory.CreateSnackRequest
	3, // 3: snackinventory.SnackInventory.ListSnacks:input_type -> snackinventory.ListSnacksRequest
	5, // 4: snackinventory.SnackInventory.DeleteSnack:input_type -> snackinventory.DeleteSnackRequest
	2, // 5: snackinventory.SnackInventory.CreateSnack:output_type -> snackinventory.CreateSnackResponse
	4, // 6: snackinventory.SnackInventory.ListSnacks:output_type -> snackinventory.ListSnacksResponse
	6, // 7: snackinventory.SnackInventory.DeleteSnack:output_type -> snackinventory.DeleteSnackResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_snackinventory_proto_init() }
func file_snackinventory_proto_init() {
	if File_snackinventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snackinventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snack); i {
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
		file_snackinventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnackRequest); i {
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
		file_snackinventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnackResponse); i {
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
		file_snackinventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnacksRequest); i {
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
		file_snackinventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnacksResponse); i {
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
		file_snackinventory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnackRequest); i {
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
		file_snackinventory_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnackResponse); i {
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
			RawDescriptor: file_snackinventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snackinventory_proto_goTypes,
		DependencyIndexes: file_snackinventory_proto_depIdxs,
		MessageInfos:      file_snackinventory_proto_msgTypes,
	}.Build()
	File_snackinventory_proto = out.File
	file_snackinventory_proto_rawDesc = nil
	file_snackinventory_proto_goTypes = nil
	file_snackinventory_proto_depIdxs = nil
}
