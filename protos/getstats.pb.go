// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protos/models/getstats.proto

package protos

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

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_models_getstats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_models_getstats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_protos_models_getstats_proto_rawDescGZIP(), []int{0}
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Blocked  bool   `protobuf:"varint,2,opt,name=blocked,proto3" json:"blocked,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_models_getstats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_models_getstats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_protos_models_getstats_proto_rawDescGZIP(), []int{1}
}

func (x *Stats) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Stats) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*Stats `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_models_getstats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_models_getstats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_protos_models_getstats_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatsResponse) GetStats() []*Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_protos_models_getstats_proto protoreflect.FileDescriptor

var file_protos_models_getstats_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x67, 0x65, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x67, 0x65, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_models_getstats_proto_rawDescOnce sync.Once
	file_protos_models_getstats_proto_rawDescData = file_protos_models_getstats_proto_rawDesc
)

func file_protos_models_getstats_proto_rawDescGZIP() []byte {
	file_protos_models_getstats_proto_rawDescOnce.Do(func() {
		file_protos_models_getstats_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_models_getstats_proto_rawDescData)
	})
	return file_protos_models_getstats_proto_rawDescData
}

var file_protos_models_getstats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_models_getstats_proto_goTypes = []interface{}{
	(*GetStatsRequest)(nil),  // 0: getstats.GetStatsRequest
	(*Stats)(nil),            // 1: getstats.Stats
	(*GetStatsResponse)(nil), // 2: getstats.GetStatsResponse
}
var file_protos_models_getstats_proto_depIdxs = []int32{
	1, // 0: getstats.GetStatsResponse.stats:type_name -> getstats.Stats
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_models_getstats_proto_init() }
func file_protos_models_getstats_proto_init() {
	if File_protos_models_getstats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_models_getstats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsRequest); i {
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
		file_protos_models_getstats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_protos_models_getstats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
			RawDescriptor: file_protos_models_getstats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_models_getstats_proto_goTypes,
		DependencyIndexes: file_protos_models_getstats_proto_depIdxs,
		MessageInfos:      file_protos_models_getstats_proto_msgTypes,
	}.Build()
	File_protos_models_getstats_proto = out.File
	file_protos_models_getstats_proto_rawDesc = nil
	file_protos_models_getstats_proto_goTypes = nil
	file_protos_models_getstats_proto_depIdxs = nil
}
