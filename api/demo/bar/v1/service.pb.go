// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: demo/bar/v1/service.proto

package barv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDrinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDrinkRequest) Reset() {
	*x = GetDrinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bar_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrinkRequest) ProtoMessage() {}

func (x *GetDrinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bar_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrinkRequest.ProtoReflect.Descriptor instead.
func (*GetDrinkRequest) Descriptor() ([]byte, []int) {
	return file_demo_bar_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDrinkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDrinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetDrinkResponse) Reset() {
	*x = GetDrinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_bar_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrinkResponse) ProtoMessage() {}

func (x *GetDrinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_bar_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrinkResponse.ProtoReflect.Descriptor instead.
func (*GetDrinkResponse) Descriptor() ([]byte, []int) {
	return file_demo_bar_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDrinkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_demo_bar_v1_service_proto protoreflect.FileDescriptor

var file_demo_bar_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xa6, 0x02, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x97, 0x02, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x12, 0x1c,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcd, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x62, 0x61, 0x72, 0x2f, 0x64, 0x72, 0x69, 0x6e, 0x6b,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x92, 0x41, 0xb0, 0x01, 0x0a, 0x05, 0x44, 0x72, 0x69,
	0x6e, 0x6b, 0x12, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x1a,
	0x13, 0x4c, 0x65, 0x74, 0x73, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x67, 0x65, 0x74, 0x61, 0x20, 0x64,
	0x72, 0x69, 0x6e, 0x6b, 0x2a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x72, 0x69, 0x6e, 0x6b, 0x4a,
	0x1d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4a, 0x3a,
	0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x33, 0x0a, 0x31, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x20, 0x64, 0x72, 0x69, 0x6e, 0x6b, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x4a, 0x1f, 0x0a, 0x03, 0x35, 0x30,
	0x30, 0x12, 0x18, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x42, 0x87, 0x01, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x77,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x62,
	0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x72, 0x76, 0x31, 0x92, 0x41, 0x47, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x52, 0x33, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x2c, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_bar_v1_service_proto_rawDescOnce sync.Once
	file_demo_bar_v1_service_proto_rawDescData = file_demo_bar_v1_service_proto_rawDesc
)

func file_demo_bar_v1_service_proto_rawDescGZIP() []byte {
	file_demo_bar_v1_service_proto_rawDescOnce.Do(func() {
		file_demo_bar_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_bar_v1_service_proto_rawDescData)
	})
	return file_demo_bar_v1_service_proto_rawDescData
}

var file_demo_bar_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_bar_v1_service_proto_goTypes = []interface{}{
	(*GetDrinkRequest)(nil),  // 0: demo.bar.v1.GetDrinkRequest
	(*GetDrinkResponse)(nil), // 1: demo.bar.v1.GetDrinkResponse
}
var file_demo_bar_v1_service_proto_depIdxs = []int32{
	0, // 0: demo.bar.v1.BarService.GetDrink:input_type -> demo.bar.v1.GetDrinkRequest
	1, // 1: demo.bar.v1.BarService.GetDrink:output_type -> demo.bar.v1.GetDrinkResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_bar_v1_service_proto_init() }
func file_demo_bar_v1_service_proto_init() {
	if File_demo_bar_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_bar_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrinkRequest); i {
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
		file_demo_bar_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrinkResponse); i {
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
			RawDescriptor: file_demo_bar_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_bar_v1_service_proto_goTypes,
		DependencyIndexes: file_demo_bar_v1_service_proto_depIdxs,
		MessageInfos:      file_demo_bar_v1_service_proto_msgTypes,
	}.Build()
	File_demo_bar_v1_service_proto = out.File
	file_demo_bar_v1_service_proto_rawDesc = nil
	file_demo_bar_v1_service_proto_goTypes = nil
	file_demo_bar_v1_service_proto_depIdxs = nil
}