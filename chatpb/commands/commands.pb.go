// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/chat/commands.proto

package commands

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

// User wants to create a new chat
type Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatName string `protobuf:"bytes,1,opt,name=chat_name,json=chatName,proto3" json:"chat_name,omitempty"`
}

func (x *Create) Reset() {
	*x = Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Create) ProtoMessage() {}

func (x *Create) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Create.ProtoReflect.Descriptor instead.
func (*Create) Descriptor() ([]byte, []int) {
	return file_protos_chat_commands_proto_rawDescGZIP(), []int{0}
}

func (x *Create) GetChatName() string {
	if x != nil {
		return x.ChatName
	}
	return ""
}

// User wants to destroy an existing chat
type Destroy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatName string `protobuf:"bytes,1,opt,name=chat_name,json=chatName,proto3" json:"chat_name,omitempty"`
}

func (x *Destroy) Reset() {
	*x = Destroy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destroy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destroy) ProtoMessage() {}

func (x *Destroy) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destroy.ProtoReflect.Descriptor instead.
func (*Destroy) Descriptor() ([]byte, []int) {
	return file_protos_chat_commands_proto_rawDescGZIP(), []int{1}
}

func (x *Destroy) GetChatName() string {
	if x != nil {
		return x.ChatName
	}
	return ""
}

// User wants to join a chat
type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatName string `protobuf:"bytes,1,opt,name=chat_name,json=chatName,proto3" json:"chat_name,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_protos_chat_commands_proto_rawDescGZIP(), []int{2}
}

func (x *Join) GetChatName() string {
	if x != nil {
		return x.ChatName
	}
	return ""
}

// User wants to leave a chat
type Unjoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Unjoin) Reset() {
	*x = Unjoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_commands_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unjoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unjoin) ProtoMessage() {}

func (x *Unjoin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_commands_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unjoin.ProtoReflect.Descriptor instead.
func (*Unjoin) Descriptor() ([]byte, []int) {
	return file_protos_chat_commands_proto_rawDescGZIP(), []int{3}
}

func (x *Unjoin) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// User wants to send a message to a chat
type Say struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Say) Reset() {
	*x = Say{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_commands_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Say) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Say) ProtoMessage() {}

func (x *Say) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_commands_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Say.ProtoReflect.Descriptor instead.
func (*Say) Descriptor() ([]byte, []int) {
	return file_protos_chat_commands_proto_rawDescGZIP(), []int{4}
}

func (x *Say) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_chat_commands_proto protoreflect.FileDescriptor

var file_protos_chat_commands_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x55, 0x6e,
	0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f,
	0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x72,
	0x6f, 0x6e, 0x53, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_chat_commands_proto_rawDescOnce sync.Once
	file_protos_chat_commands_proto_rawDescData = file_protos_chat_commands_proto_rawDesc
)

func file_protos_chat_commands_proto_rawDescGZIP() []byte {
	file_protos_chat_commands_proto_rawDescOnce.Do(func() {
		file_protos_chat_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_chat_commands_proto_rawDescData)
	})
	return file_protos_chat_commands_proto_rawDescData
}

var file_protos_chat_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_chat_commands_proto_goTypes = []interface{}{
	(*Create)(nil),  // 0: commands.Create
	(*Destroy)(nil), // 1: commands.Destroy
	(*Join)(nil),    // 2: commands.Join
	(*Unjoin)(nil),  // 3: commands.Unjoin
	(*Say)(nil),     // 4: commands.Say
}
var file_protos_chat_commands_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_chat_commands_proto_init() }
func file_protos_chat_commands_proto_init() {
	if File_protos_chat_commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_chat_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Create); i {
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
		file_protos_chat_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destroy); i {
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
		file_protos_chat_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
		file_protos_chat_commands_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unjoin); i {
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
		file_protos_chat_commands_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Say); i {
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
			RawDescriptor: file_protos_chat_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_chat_commands_proto_goTypes,
		DependencyIndexes: file_protos_chat_commands_proto_depIdxs,
		MessageInfos:      file_protos_chat_commands_proto_msgTypes,
	}.Build()
	File_protos_chat_commands_proto = out.File
	file_protos_chat_commands_proto_rawDesc = nil
	file_protos_chat_commands_proto_goTypes = nil
	file_protos_chat_commands_proto_depIdxs = nil
}