// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/example/pb/request.proto

package example

import (
	proto "github.com/golang/protobuf/proto"
	page "github.com/infraboard/mcube/pb/page"
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

// CreateBookRequest 创建Book请求
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_example_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_example_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_pkg_example_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryBookRequest 查询Book请求
type QueryBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryBookRequest) Reset() {
	*x = QueryBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_example_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBookRequest) ProtoMessage() {}

func (x *QueryBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_example_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBookRequest.ProtoReflect.Descriptor instead.
func (*QueryBookRequest) Descriptor() ([]byte, []int) {
	return file_pkg_example_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryBookRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pkg_example_pb_request_proto protoreflect.FileDescriptor

var file_pkg_example_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_example_pb_request_proto_rawDescOnce sync.Once
	file_pkg_example_pb_request_proto_rawDescData = file_pkg_example_pb_request_proto_rawDesc
)

func file_pkg_example_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_example_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_example_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_example_pb_request_proto_rawDescData)
	})
	return file_pkg_example_pb_request_proto_rawDescData
}

var file_pkg_example_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_example_pb_request_proto_goTypes = []interface{}{
	(*CreateBookRequest)(nil), // 0: demo.example.CreateBookRequest
	(*QueryBookRequest)(nil),  // 1: demo.example.QueryBookRequest
	(*page.PageRequest)(nil),  // 2: page.PageRequest
}
var file_pkg_example_pb_request_proto_depIdxs = []int32{
	2, // 0: demo.example.QueryBookRequest.page:type_name -> page.PageRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_example_pb_request_proto_init() }
func file_pkg_example_pb_request_proto_init() {
	if File_pkg_example_pb_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_example_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_pkg_example_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBookRequest); i {
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
			RawDescriptor: file_pkg_example_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_example_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_example_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_example_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_example_pb_request_proto = out.File
	file_pkg_example_pb_request_proto_rawDesc = nil
	file_pkg_example_pb_request_proto_goTypes = nil
	file_pkg_example_pb_request_proto_depIdxs = nil
}
