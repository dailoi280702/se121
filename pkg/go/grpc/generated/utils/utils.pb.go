// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: utils.proto

package utils

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

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query       *string `protobuf:"bytes,1,opt,name=query,proto3,oneof" json:"query,omitempty"`
	Orderby     *string `protobuf:"bytes,2,opt,name=orderby,proto3,oneof" json:"orderby,omitempty"`
	StartAt     *int32  `protobuf:"varint,3,opt,name=startAt,proto3,oneof" json:"startAt,omitempty"`
	Limit       *int32  `protobuf:"varint,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	IsAscending *bool   `protobuf:"varint,5,opt,name=isAscending,proto3,oneof" json:"isAscending,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_utils_proto_rawDescGZIP(), []int{0}
}

func (x *SearchReq) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *SearchReq) GetOrderby() string {
	if x != nil && x.Orderby != nil {
		return *x.Orderby
	}
	return ""
}

func (x *SearchReq) GetStartAt() int32 {
	if x != nil && x.StartAt != nil {
		return *x.StartAt
	}
	return 0
}

func (x *SearchReq) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *SearchReq) GetIsAscending() bool {
	if x != nil && x.IsAscending != nil {
		return *x.IsAscending
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_utils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_utils_proto_rawDescGZIP(), []int{1}
}

var File_utils_proto protoreflect.FileDescriptor

var file_utils_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75,
	0x74, 0x69, 0x6c, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x73, 0x63, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x0b, 0x69,
	0x73, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x62, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x73,
	0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x69, 0x6c, 0x6f, 0x69, 0x32, 0x38, 0x30, 0x37, 0x30, 0x32, 0x2f, 0x73, 0x65,
	0x31, 0x32, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utils_proto_rawDescOnce sync.Once
	file_utils_proto_rawDescData = file_utils_proto_rawDesc
)

func file_utils_proto_rawDescGZIP() []byte {
	file_utils_proto_rawDescOnce.Do(func() {
		file_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_utils_proto_rawDescData)
	})
	return file_utils_proto_rawDescData
}

var file_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_utils_proto_goTypes = []interface{}{
	(*SearchReq)(nil), // 0: utils.SearchReq
	(*Empty)(nil),     // 1: utils.Empty
}
var file_utils_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_utils_proto_init() }
func file_utils_proto_init() {
	if File_utils_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_utils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
	file_utils_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_utils_proto_goTypes,
		DependencyIndexes: file_utils_proto_depIdxs,
		MessageInfos:      file_utils_proto_msgTypes,
	}.Build()
	File_utils_proto = out.File
	file_utils_proto_rawDesc = nil
	file_utils_proto_goTypes = nil
	file_utils_proto_depIdxs = nil
}
