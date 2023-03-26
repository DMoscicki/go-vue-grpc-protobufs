// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: answer.proto

package protogen

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

type GetTodos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTodos) Reset() {
	*x = GetTodos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodos) ProtoMessage() {}

func (x *GetTodos) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodos.ProtoReflect.Descriptor instead.
func (*GetTodos) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{0}
}

type GetRepos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRepos) Reset() {
	*x = GetRepos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepos) ProtoMessage() {}

func (x *GetRepos) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepos.ProtoReflect.Descriptor instead.
func (*GetRepos) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{1}
}

type AddTodos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *AddTodos) Reset() {
	*x = AddTodos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodos) ProtoMessage() {}

func (x *AddTodos) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodos.ProtoReflect.Descriptor instead.
func (*AddTodos) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{2}
}

func (x *AddTodos) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type DeleteTodos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodos) Reset() {
	*x = DeleteTodos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodos) ProtoMessage() {}

func (x *DeleteTodos) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodos.ProtoReflect.Descriptor instead.
func (*DeleteTodos) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTodos) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TodoObject) Reset() {
	*x = TodoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoObject) ProtoMessage() {}

func (x *TodoObject) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoObject.ProtoReflect.Descriptor instead.
func (*TodoObject) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{4}
}

func (x *TodoObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoObject) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*TodoObject `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{5}
}

func (x *TodoResponse) GetTodos() []*TodoObject {
	if x != nil {
		return x.Todos
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{7}
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LanguageLib struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stars int64  `protobuf:"varint,2,opt,name=stars,proto3" json:"stars,omitempty"`
}

func (x *LanguageLib) Reset() {
	*x = LanguageLib{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageLib) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageLib) ProtoMessage() {}

func (x *LanguageLib) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageLib.ProtoReflect.Descriptor instead.
func (*LanguageLib) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{8}
}

func (x *LanguageLib) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LanguageLib) GetStars() int64 {
	if x != nil {
		return x.Stars
	}
	return 0
}

type LanguageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Libraries []*LanguageLib `protobuf:"bytes,1,rep,name=libraries,proto3" json:"libraries,omitempty"`
}

func (x *LanguageResponse) Reset() {
	*x = LanguageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_answer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageResponse) ProtoMessage() {}

func (x *LanguageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_answer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageResponse.ProtoReflect.Descriptor instead.
func (*LanguageResponse) Descriptor() ([]byte, []int) {
	return file_answer_proto_rawDescGZIP(), []int{9}
}

func (x *LanguageResponse) GetLibraries() []*LanguageLib {
	if x != nil {
		return x.Libraries
	}
	return nil
}

var File_answer_proto protoreflect.FileDescriptor

var file_answer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x22, 0x1e, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x4a, 0x08, 0x08, 0x02, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22, 0x30, 0x0a, 0x0a, 0x54,
	0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x3c, 0x0a,
	0x0c, 0x74, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x2a, 0x0a, 0x0e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0b, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x73,
	0x22, 0x49, 0x0a, 0x10, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x62,
	0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x32, 0xd2, 0x02, 0x0a, 0x0b,
	0x74, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x67,
	0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x6c, 0x69, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x08, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_answer_proto_rawDescOnce sync.Once
	file_answer_proto_rawDescData = file_answer_proto_rawDesc
)

func file_answer_proto_rawDescGZIP() []byte {
	file_answer_proto_rawDescOnce.Do(func() {
		file_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_answer_proto_rawDescData)
	})
	return file_answer_proto_rawDescData
}

var file_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_answer_proto_goTypes = []interface{}{
	(*GetTodos)(nil),         // 0: protofiles.GetTodos
	(*GetRepos)(nil),         // 1: protofiles.GetRepos
	(*AddTodos)(nil),         // 2: protofiles.AddTodos
	(*DeleteTodos)(nil),      // 3: protofiles.DeleteTodos
	(*TodoObject)(nil),       // 4: protofiles.TodoObject
	(*TodoResponse)(nil),     // 5: protofiles.todoResponse
	(*DeleteResponse)(nil),   // 6: protofiles.deleteResponse
	(*Language)(nil),         // 7: protofiles.Language
	(*LanguageLib)(nil),      // 8: protofiles.LanguageLib
	(*LanguageResponse)(nil), // 9: protofiles.LanguageResponse
}
var file_answer_proto_depIdxs = []int32{
	4, // 0: protofiles.todoResponse.todos:type_name -> protofiles.TodoObject
	8, // 1: protofiles.LanguageResponse.libraries:type_name -> protofiles.LanguageLib
	2, // 2: protofiles.todoService.addTodo:input_type -> protofiles.AddTodos
	3, // 3: protofiles.todoService.deleteTodo:input_type -> protofiles.DeleteTodos
	0, // 4: protofiles.todoService.getTodos:input_type -> protofiles.GetTodos
	7, // 5: protofiles.todoService.libResponse:input_type -> protofiles.Language
	1, // 6: protofiles.todoService.getRepos:input_type -> protofiles.GetRepos
	4, // 7: protofiles.todoService.addTodo:output_type -> protofiles.TodoObject
	6, // 8: protofiles.todoService.deleteTodo:output_type -> protofiles.deleteResponse
	5, // 9: protofiles.todoService.getTodos:output_type -> protofiles.todoResponse
	9, // 10: protofiles.todoService.libResponse:output_type -> protofiles.LanguageResponse
	9, // 11: protofiles.todoService.getRepos:output_type -> protofiles.LanguageResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_answer_proto_init() }
func file_answer_proto_init() {
	if File_answer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodos); i {
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
		file_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepos); i {
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
		file_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodos); i {
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
		file_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodos); i {
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
		file_answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoObject); i {
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
		file_answer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_answer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_answer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_answer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageLib); i {
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
		file_answer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageResponse); i {
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
			RawDescriptor: file_answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_answer_proto_goTypes,
		DependencyIndexes: file_answer_proto_depIdxs,
		MessageInfos:      file_answer_proto_msgTypes,
	}.Build()
	File_answer_proto = out.File
	file_answer_proto_rawDesc = nil
	file_answer_proto_goTypes = nil
	file_answer_proto_depIdxs = nil
}
