// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protos/getauthtoken/getauthtoken.proto

package getauthtoken

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

type GetAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature     string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	WalletAddress string `protobuf:"bytes,2,opt,name=walletAddress,proto3" json:"walletAddress,omitempty"`
}

func (x *GetAuthTokenRequest) Reset() {
	*x = GetAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_getauthtoken_getauthtoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthTokenRequest) ProtoMessage() {}

func (x *GetAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_getauthtoken_getauthtoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_protos_getauthtoken_getauthtoken_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthTokenRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *GetAuthTokenRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

type GetAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAuthTokenResponse) Reset() {
	*x = GetAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_getauthtoken_getauthtoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthTokenResponse) ProtoMessage() {}

func (x *GetAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_getauthtoken_getauthtoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_protos_getauthtoken_getauthtoken_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_protos_getauthtoken_getauthtoken_proto protoreflect.FileDescriptor

var file_protos_getauthtoken_getauthtoken_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x75, 0x74, 0x68,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x75, 0x74, 0x68, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x65, 0x74, 0x61, 0x75, 0x74,
	0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x29, 0x5a, 0x27, 0x64, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x61, 0x75, 0x74, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_getauthtoken_getauthtoken_proto_rawDescOnce sync.Once
	file_protos_getauthtoken_getauthtoken_proto_rawDescData = file_protos_getauthtoken_getauthtoken_proto_rawDesc
)

func file_protos_getauthtoken_getauthtoken_proto_rawDescGZIP() []byte {
	file_protos_getauthtoken_getauthtoken_proto_rawDescOnce.Do(func() {
		file_protos_getauthtoken_getauthtoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_getauthtoken_getauthtoken_proto_rawDescData)
	})
	return file_protos_getauthtoken_getauthtoken_proto_rawDescData
}

var file_protos_getauthtoken_getauthtoken_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_getauthtoken_getauthtoken_proto_goTypes = []interface{}{
	(*GetAuthTokenRequest)(nil),  // 0: getauthtoken.GetAuthTokenRequest
	(*GetAuthTokenResponse)(nil), // 1: getauthtoken.GetAuthTokenResponse
}
var file_protos_getauthtoken_getauthtoken_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_getauthtoken_getauthtoken_proto_init() }
func file_protos_getauthtoken_getauthtoken_proto_init() {
	if File_protos_getauthtoken_getauthtoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_getauthtoken_getauthtoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthTokenRequest); i {
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
		file_protos_getauthtoken_getauthtoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthTokenResponse); i {
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
			RawDescriptor: file_protos_getauthtoken_getauthtoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_getauthtoken_getauthtoken_proto_goTypes,
		DependencyIndexes: file_protos_getauthtoken_getauthtoken_proto_depIdxs,
		MessageInfos:      file_protos_getauthtoken_getauthtoken_proto_msgTypes,
	}.Build()
	File_protos_getauthtoken_getauthtoken_proto = out.File
	file_protos_getauthtoken_getauthtoken_proto_rawDesc = nil
	file_protos_getauthtoken_getauthtoken_proto_goTypes = nil
	file_protos_getauthtoken_getauthtoken_proto_depIdxs = nil
}