// 定义我们接口的版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: data.proto

// 定义包名称

package data

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

type CreateTempInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CreateTempInfoReq) Reset() {
	*x = CreateTempInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTempInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTempInfoReq) ProtoMessage() {}

func (x *CreateTempInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTempInfoReq.ProtoReflect.Descriptor instead.
func (*CreateTempInfoReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTempInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTempInfoReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type CreateTempInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateTempInfoResp) Reset() {
	*x = CreateTempInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTempInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTempInfoResp) ProtoMessage() {}

func (x *CreateTempInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTempInfoResp.ProtoReflect.Descriptor instead.
func (*CreateTempInfoResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTempInfoResp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UploadTempFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadTempFileReq) Reset() {
	*x = UploadTempFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadTempFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTempFileReq) ProtoMessage() {}

func (x *UploadTempFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTempFileReq.ProtoReflect.Descriptor instead.
func (*UploadTempFileReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *UploadTempFileReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UploadTempFileReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadTempFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UploadTempFileResp) Reset() {
	*x = UploadTempFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadTempFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTempFileResp) ProtoMessage() {}

func (x *UploadTempFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTempFileResp.ProtoReflect.Descriptor instead.
func (*UploadTempFileResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *UploadTempFileResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CommitTempFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CommitTempFileReq) Reset() {
	*x = CommitTempFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTempFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTempFileReq) ProtoMessage() {}

func (x *CommitTempFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTempFileReq.ProtoReflect.Descriptor instead.
func (*CommitTempFileReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *CommitTempFileReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CommitTempFileReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type CommitTempFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CommitTempFileResp) Reset() {
	*x = CommitTempFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTempFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTempFileResp) ProtoMessage() {}

func (x *CommitTempFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTempFileResp.ProtoReflect.Descriptor instead.
func (*CommitTempFileResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{5}
}

func (x *CommitTempFileResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CommonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CommonReq) Reset() {
	*x = CommonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReq) ProtoMessage() {}

func (x *CommonReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReq.ProtoReflect.Descriptor instead.
func (*CommonReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{6}
}

func (x *CommonReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteTempFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteTempFileResp) Reset() {
	*x = DeleteTempFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTempFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTempFileResp) ProtoMessage() {}

func (x *DeleteTempFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTempFileResp.ProtoReflect.Descriptor instead.
func (*DeleteTempFileResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTempFileResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetTempFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTempFileResp) Reset() {
	*x = GetTempFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTempFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTempFileResp) ProtoMessage() {}

func (x *GetTempFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTempFileResp.ProtoReflect.Descriptor instead.
func (*GetTempFileResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{8}
}

func (x *GetTempFileResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HeadTempFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ContentLength int64 `protobuf:"varint,2,opt,name=contentLength,proto3" json:"contentLength,omitempty"`
}

func (x *HeadTempFileResp) Reset() {
	*x = HeadTempFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadTempFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadTempFileResp) ProtoMessage() {}

func (x *HeadTempFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadTempFileResp.ProtoReflect.Descriptor instead.
func (*HeadTempFileResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{9}
}

func (x *HeadTempFileResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HeadTempFileResp) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

type GetMatterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamePrefix string `protobuf:"bytes,1,opt,name=namePrefix,proto3" json:"namePrefix,omitempty"`
}

func (x *GetMatterReq) Reset() {
	*x = GetMatterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatterReq) ProtoMessage() {}

func (x *GetMatterReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatterReq.ProtoReflect.Descriptor instead.
func (*GetMatterReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{10}
}

func (x *GetMatterReq) GetNamePrefix() string {
	if x != nil {
		return x.NamePrefix
	}
	return ""
}

type GetMatterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMatterResp) Reset() {
	*x = GetMatterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatterResp) ProtoMessage() {}

func (x *GetMatterResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatterResp.ProtoReflect.Descriptor instead.
func (*GetMatterResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{11}
}

func (x *GetMatterResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CheckTempFileHashReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Hash   string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *CheckTempFileHashReq) Reset() {
	*x = CheckTempFileHashReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTempFileHashReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTempFileHashReq) ProtoMessage() {}

func (x *CheckTempFileHashReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTempFileHashReq.ProtoReflect.Descriptor instead.
func (*CheckTempFileHashReq) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{12}
}

func (x *CheckTempFileHashReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CheckTempFileHashReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *CheckTempFileHashReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type CheckTempFileHashResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CheckTempFileHashResp) Reset() {
	*x = CheckTempFileHashResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTempFileHashResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTempFileHashResp) ProtoMessage() {}

func (x *CheckTempFileHashResp) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTempFileHashResp.ProtoReflect.Descriptor instead.
func (*CheckTempFileHashResp) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{13}
}

func (x *CheckTempFileHashResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x28,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54,
	0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x2c, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x25, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x64, 0x54, 0x65, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x14, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xfc, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x43, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54,
	0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x28, 0x01, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x4a, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x6d,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_data_proto_goTypes = []interface{}{
	(*CreateTempInfoReq)(nil),     // 0: api.CreateTempInfoReq
	(*CreateTempInfoResp)(nil),    // 1: api.CreateTempInfoResp
	(*UploadTempFileReq)(nil),     // 2: api.UploadTempFileReq
	(*UploadTempFileResp)(nil),    // 3: api.UploadTempFileResp
	(*CommitTempFileReq)(nil),     // 4: api.CommitTempFileReq
	(*CommitTempFileResp)(nil),    // 5: api.CommitTempFileResp
	(*CommonReq)(nil),             // 6: api.CommonReq
	(*DeleteTempFileResp)(nil),    // 7: api.DeleteTempFileResp
	(*GetTempFileResp)(nil),       // 8: api.GetTempFileResp
	(*HeadTempFileResp)(nil),      // 9: api.HeadTempFileResp
	(*GetMatterReq)(nil),          // 10: api.GetMatterReq
	(*GetMatterResp)(nil),         // 11: api.GetMatterResp
	(*CheckTempFileHashReq)(nil),  // 12: api.CheckTempFileHashReq
	(*CheckTempFileHashResp)(nil), // 13: api.CheckTempFileHashResp
}
var file_data_proto_depIdxs = []int32{
	0,  // 0: api.Data.CreateTempInfo:input_type -> api.CreateTempInfoReq
	2,  // 1: api.Data.UploadTempFile:input_type -> api.UploadTempFileReq
	4,  // 2: api.Data.CommitTempFile:input_type -> api.CommitTempFileReq
	6,  // 3: api.Data.DeleteTempFile:input_type -> api.CommonReq
	6,  // 4: api.Data.GetTempFile:input_type -> api.CommonReq
	6,  // 5: api.Data.HeadTempFile:input_type -> api.CommonReq
	12, // 6: api.Data.CheckTempFileHash:input_type -> api.CheckTempFileHashReq
	10, // 7: api.Data.GetMatter:input_type -> api.GetMatterReq
	1,  // 8: api.Data.CreateTempInfo:output_type -> api.CreateTempInfoResp
	3,  // 9: api.Data.UploadTempFile:output_type -> api.UploadTempFileResp
	5,  // 10: api.Data.CommitTempFile:output_type -> api.CommitTempFileResp
	7,  // 11: api.Data.DeleteTempFile:output_type -> api.DeleteTempFileResp
	8,  // 12: api.Data.GetTempFile:output_type -> api.GetTempFileResp
	9,  // 13: api.Data.HeadTempFile:output_type -> api.HeadTempFileResp
	13, // 14: api.Data.CheckTempFileHash:output_type -> api.CheckTempFileHashResp
	11, // 15: api.Data.GetMatter:output_type -> api.GetMatterResp
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTempInfoReq); i {
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
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTempInfoResp); i {
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
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadTempFileReq); i {
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
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadTempFileResp); i {
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
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTempFileReq); i {
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
		file_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTempFileResp); i {
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
		file_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReq); i {
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
		file_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTempFileResp); i {
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
		file_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTempFileResp); i {
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
		file_data_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadTempFileResp); i {
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
		file_data_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatterReq); i {
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
		file_data_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatterResp); i {
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
		file_data_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTempFileHashReq); i {
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
		file_data_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTempFileHashResp); i {
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
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
