// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v1/key.proto

package v1

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

type KeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey  string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Status     string `protobuf:"bytes,99,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *KeyResponse) Reset() {
	*x = KeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyResponse) ProtoMessage() {}

func (x *KeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyResponse.ProtoReflect.Descriptor instead.
func (*KeyResponse) Descriptor() ([]byte, []int) {
	return file_v1_key_proto_rawDescGZIP(), []int{0}
}

func (x *KeyResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *KeyResponse) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *KeyResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DecodeKey string `protobuf:"bytes,2,opt,name=decode_key,json=decodeKey,proto3" json:"decode_key,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_v1_key_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *KeyRequest) GetDecodeKey() string {
	if x != nil {
		return x.DecodeKey
	}
	return ""
}

var File_v1_key_proto protoreflect.FileDescriptor

var file_v1_key_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x63, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x44, 0x0a,
	0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x4b, 0x65, 0x79, 0x32, 0xab, 0x01, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x2e, 0x6b,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6b,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_key_proto_rawDescOnce sync.Once
	file_v1_key_proto_rawDescData = file_v1_key_proto_rawDesc
)

func file_v1_key_proto_rawDescGZIP() []byte {
	file_v1_key_proto_rawDescOnce.Do(func() {
		file_v1_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_key_proto_rawDescData)
	})
	return file_v1_key_proto_rawDescData
}

var file_v1_key_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_key_proto_goTypes = []interface{}{
	(*KeyResponse)(nil), // 0: key.v1.KeyResponse
	(*KeyRequest)(nil),  // 1: key.v1.KeyRequest
}
var file_v1_key_proto_depIdxs = []int32{
	1, // 0: key.v1.KeyService.GetKey:input_type -> key.v1.KeyRequest
	1, // 1: key.v1.KeyService.CreateKey:input_type -> key.v1.KeyRequest
	1, // 2: key.v1.KeyService.DeleteKey:input_type -> key.v1.KeyRequest
	0, // 3: key.v1.KeyService.GetKey:output_type -> key.v1.KeyResponse
	0, // 4: key.v1.KeyService.CreateKey:output_type -> key.v1.KeyResponse
	0, // 5: key.v1.KeyService.DeleteKey:output_type -> key.v1.KeyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_key_proto_init() }
func file_v1_key_proto_init() {
	if File_v1_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyResponse); i {
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
		file_v1_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
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
			RawDescriptor: file_v1_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_key_proto_goTypes,
		DependencyIndexes: file_v1_key_proto_depIdxs,
		MessageInfos:      file_v1_key_proto_msgTypes,
	}.Build()
	File_v1_key_proto = out.File
	file_v1_key_proto_rawDesc = nil
	file_v1_key_proto_goTypes = nil
	file_v1_key_proto_depIdxs = nil
}
