// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: text_generate.proto

package text_generate

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

type GenerateReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Brand        *string `protobuf:"bytes,2,opt,name=brand,proto3,oneof" json:"brand,omitempty"`
	Series       *string `protobuf:"bytes,3,opt,name=series,proto3,oneof" json:"series,omitempty"`
	HorsePower   *int32  `protobuf:"varint,4,opt,name=horsePower,proto3,oneof" json:"horsePower,omitempty"`
	Torque       *int32  `protobuf:"varint,5,opt,name=torque,proto3,oneof" json:"torque,omitempty"`
	Transmission *string `protobuf:"bytes,6,opt,name=transmission,proto3,oneof" json:"transmission,omitempty"`
	FuelType     *string `protobuf:"bytes,7,opt,name=fuelType,proto3,oneof" json:"fuelType,omitempty"`
}

func (x *GenerateReviewReq) Reset() {
	*x = GenerateReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_generate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateReviewReq) ProtoMessage() {}

func (x *GenerateReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_text_generate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateReviewReq.ProtoReflect.Descriptor instead.
func (*GenerateReviewReq) Descriptor() ([]byte, []int) {
	return file_text_generate_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateReviewReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenerateReviewReq) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *GenerateReviewReq) GetSeries() string {
	if x != nil && x.Series != nil {
		return *x.Series
	}
	return ""
}

func (x *GenerateReviewReq) GetHorsePower() int32 {
	if x != nil && x.HorsePower != nil {
		return *x.HorsePower
	}
	return 0
}

func (x *GenerateReviewReq) GetTorque() int32 {
	if x != nil && x.Torque != nil {
		return *x.Torque
	}
	return 0
}

func (x *GenerateReviewReq) GetTransmission() string {
	if x != nil && x.Transmission != nil {
		return *x.Transmission
	}
	return ""
}

func (x *GenerateReviewReq) GetFuelType() string {
	if x != nil && x.FuelType != nil {
		return *x.FuelType
	}
	return ""
}

type GenerateBlogSummarizationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GenerateBlogSummarizationReq) Reset() {
	*x = GenerateBlogSummarizationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_generate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateBlogSummarizationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateBlogSummarizationReq) ProtoMessage() {}

func (x *GenerateBlogSummarizationReq) ProtoReflect() protoreflect.Message {
	mi := &file_text_generate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateBlogSummarizationReq.ProtoReflect.Descriptor instead.
func (*GenerateBlogSummarizationReq) Descriptor() ([]byte, []int) {
	return file_text_generate_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateBlogSummarizationReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GenerateBlogSummarizationReq) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type ResString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ResString) Reset() {
	*x = ResString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_generate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResString) ProtoMessage() {}

func (x *ResString) ProtoReflect() protoreflect.Message {
	mi := &file_text_generate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResString.ProtoReflect.Descriptor instead.
func (*ResString) Descriptor() ([]byte, []int) {
	return file_text_generate_proto_rawDescGZIP(), []int{2}
}

func (x *ResString) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_text_generate_proto protoreflect.FileDescriptor

var file_text_generate_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x22, 0xb8, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0a, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x74,
	0x6f, 0x72, 0x71, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x06, 0x74,
	0x6f, 0x72, 0x71, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x68, 0x6f, 0x72, 0x73,
	0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x6f, 0x72, 0x71, 0x75,
	0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x48, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1f, 0x0a, 0x09, 0x52, 0x65, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xca, 0x01, 0x0a, 0x13, 0x54,
	0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x62, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x6f, 0x69, 0x32, 0x38, 0x30, 0x37,
	0x30, 0x32, 0x2f, 0x73, 0x65, 0x31, 0x32, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_text_generate_proto_rawDescOnce sync.Once
	file_text_generate_proto_rawDescData = file_text_generate_proto_rawDesc
)

func file_text_generate_proto_rawDescGZIP() []byte {
	file_text_generate_proto_rawDescOnce.Do(func() {
		file_text_generate_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_generate_proto_rawDescData)
	})
	return file_text_generate_proto_rawDescData
}

var file_text_generate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_text_generate_proto_goTypes = []interface{}{
	(*GenerateReviewReq)(nil),            // 0: text_generate.GenerateReviewReq
	(*GenerateBlogSummarizationReq)(nil), // 1: text_generate.GenerateBlogSummarizationReq
	(*ResString)(nil),                    // 2: text_generate.ResString
}
var file_text_generate_proto_depIdxs = []int32{
	0, // 0: text_generate.TextGenerateService.GenerateCarReview:input_type -> text_generate.GenerateReviewReq
	1, // 1: text_generate.TextGenerateService.GenerateBlogSummarization:input_type -> text_generate.GenerateBlogSummarizationReq
	2, // 2: text_generate.TextGenerateService.GenerateCarReview:output_type -> text_generate.ResString
	2, // 3: text_generate.TextGenerateService.GenerateBlogSummarization:output_type -> text_generate.ResString
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_text_generate_proto_init() }
func file_text_generate_proto_init() {
	if File_text_generate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_generate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateReviewReq); i {
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
		file_text_generate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateBlogSummarizationReq); i {
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
		file_text_generate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResString); i {
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
	file_text_generate_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_text_generate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_text_generate_proto_goTypes,
		DependencyIndexes: file_text_generate_proto_depIdxs,
		MessageInfos:      file_text_generate_proto_msgTypes,
	}.Build()
	File_text_generate_proto = out.File
	file_text_generate_proto_rawDesc = nil
	file_text_generate_proto_goTypes = nil
	file_text_generate_proto_depIdxs = nil
}
