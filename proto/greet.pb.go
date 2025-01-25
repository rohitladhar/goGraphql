// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: proto/greet.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	mi := &file_proto_greet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{0}
}

type HelloReequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloReequest) Reset() {
	*x = HelloReequest{}
	mi := &file_proto_greet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloReequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReequest) ProtoMessage() {}

func (x *HelloReequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReequest.ProtoReflect.Descriptor instead.
func (*HelloReequest) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_proto_greet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NamesList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Names         []string               `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamesList) Reset() {
	*x = NamesList{}
	mi := &file_proto_greet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamesList) ProtoMessage() {}

func (x *NamesList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamesList.ProtoReflect.Descriptor instead.
func (*NamesList) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{3}
}

func (x *NamesList) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type MessagesList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []string               `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessagesList) Reset() {
	*x = MessagesList{}
	mi := &file_proto_greet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessagesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesList) ProtoMessage() {}

func (x *MessagesList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_greet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesList.ProtoReflect.Descriptor instead.
func (*MessagesList) Descriptor() ([]byte, []int) {
	return file_proto_greet_proto_rawDescGZIP(), []int{4}
}

func (x *MessagesList) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_greet_proto protoreflect.FileDescriptor

var file_proto_greet_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x76,
	0x65, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x23, 0x0a, 0x0d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x2a, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xd7, 0x02, 0x0a, 0x0c,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x17,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x54, 0x0a, 0x17, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x28, 0x01, 0x12, 0x5e, 0x0a, 0x1e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_greet_proto_rawDescOnce sync.Once
	file_proto_greet_proto_rawDescData []byte
)

func file_proto_greet_proto_rawDescGZIP() []byte {
	file_proto_greet_proto_rawDescOnce.Do(func() {
		file_proto_greet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_greet_proto_rawDesc), len(file_proto_greet_proto_rawDesc)))
	})
	return file_proto_greet_proto_rawDescData
}

var file_proto_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_greet_proto_goTypes = []any{
	(*NoParam)(nil),       // 0: greet_serive.NoParam
	(*HelloReequest)(nil), // 1: greet_serive.HelloReequest
	(*HelloResponse)(nil), // 2: greet_serive.HelloResponse
	(*NamesList)(nil),     // 3: greet_serive.NamesList
	(*MessagesList)(nil),  // 4: greet_serive.MessagesList
}
var file_proto_greet_proto_depIdxs = []int32{
	0, // 0: greet_serive.GreetService.SayHello:input_type -> greet_serive.NoParam
	3, // 1: greet_serive.GreetService.SayHelloServerStreaming:input_type -> greet_serive.NamesList
	2, // 2: greet_serive.GreetService.SayHelloClientStreaming:input_type -> greet_serive.HelloResponse
	2, // 3: greet_serive.GreetService.SayHelloBidirectionalStreaming:input_type -> greet_serive.HelloResponse
	2, // 4: greet_serive.GreetService.SayHello:output_type -> greet_serive.HelloResponse
	2, // 5: greet_serive.GreetService.SayHelloServerStreaming:output_type -> greet_serive.HelloResponse
	4, // 6: greet_serive.GreetService.SayHelloClientStreaming:output_type -> greet_serive.MessagesList
	2, // 7: greet_serive.GreetService.SayHelloBidirectionalStreaming:output_type -> greet_serive.HelloResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_greet_proto_init() }
func file_proto_greet_proto_init() {
	if File_proto_greet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_greet_proto_rawDesc), len(file_proto_greet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_greet_proto_goTypes,
		DependencyIndexes: file_proto_greet_proto_depIdxs,
		MessageInfos:      file_proto_greet_proto_msgTypes,
	}.Build()
	File_proto_greet_proto = out.File
	file_proto_greet_proto_goTypes = nil
	file_proto_greet_proto_depIdxs = nil
}
