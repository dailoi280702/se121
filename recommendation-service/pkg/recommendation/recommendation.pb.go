// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: recommendation.proto

package recommendation

import (
	blog "github.com/dailoi280702/se121/blog-service/pkg/blog"
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

type GetRelatedBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId       int32 `protobuf:"varint,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	NumberOfBlog int32 `protobuf:"varint,2,opt,name=numberOfBlog,proto3" json:"numberOfBlog,omitempty"`
}

func (x *GetRelatedBlogReq) Reset() {
	*x = GetRelatedBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelatedBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelatedBlogReq) ProtoMessage() {}

func (x *GetRelatedBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelatedBlogReq.ProtoReflect.Descriptor instead.
func (*GetRelatedBlogReq) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{0}
}

func (x *GetRelatedBlogReq) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *GetRelatedBlogReq) GetNumberOfBlog() int32 {
	if x != nil {
		return x.NumberOfBlog
	}
	return 0
}

type GetRelatedBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blogs []*blog.Blog `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
}

func (x *GetRelatedBlogRes) Reset() {
	*x = GetRelatedBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelatedBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelatedBlogRes) ProtoMessage() {}

func (x *GetRelatedBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelatedBlogRes.ProtoReflect.Descriptor instead.
func (*GetRelatedBlogRes) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{1}
}

func (x *GetRelatedBlogRes) GetBlogs() []*blog.Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type GetUserRecommendedBlogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetUserRecommendedBlogsReq) Reset() {
	*x = GetUserRecommendedBlogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRecommendedBlogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRecommendedBlogsReq) ProtoMessage() {}

func (x *GetUserRecommendedBlogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRecommendedBlogsReq.ProtoReflect.Descriptor instead.
func (*GetUserRecommendedBlogsReq) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRecommendedBlogsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserRecommendedBlogsReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_recommendation_proto protoreflect.FileDescriptor

var file_recommendation_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x1a, 0x0a, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x42, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x35, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x20,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x22, 0x4a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0xa5, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x6f, 0x69, 0x32, 0x38, 0x30, 0x37, 0x30, 0x32, 0x2f,
	0x73, 0x65, 0x31, 0x32, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommendation_proto_rawDescOnce sync.Once
	file_recommendation_proto_rawDescData = file_recommendation_proto_rawDesc
)

func file_recommendation_proto_rawDescGZIP() []byte {
	file_recommendation_proto_rawDescOnce.Do(func() {
		file_recommendation_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommendation_proto_rawDescData)
	})
	return file_recommendation_proto_rawDescData
}

var file_recommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_recommendation_proto_goTypes = []interface{}{
	(*GetRelatedBlogReq)(nil),          // 0: blog.GetRelatedBlogReq
	(*GetRelatedBlogRes)(nil),          // 1: blog.GetRelatedBlogRes
	(*GetUserRecommendedBlogsReq)(nil), // 2: blog.GetUserRecommendedBlogsReq
	(*blog.Blog)(nil),                  // 3: blog.Blog
	(*blog.Blogs)(nil),                 // 4: blog.Blogs
}
var file_recommendation_proto_depIdxs = []int32{
	3, // 0: blog.GetRelatedBlogRes.blogs:type_name -> blog.Blog
	0, // 1: blog.RecommendationService.GetRelatedBlog:input_type -> blog.GetRelatedBlogReq
	2, // 2: blog.RecommendationService.GetUserRecommendedBlogs:input_type -> blog.GetUserRecommendedBlogsReq
	1, // 3: blog.RecommendationService.GetRelatedBlog:output_type -> blog.GetRelatedBlogRes
	4, // 4: blog.RecommendationService.GetUserRecommendedBlogs:output_type -> blog.Blogs
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recommendation_proto_init() }
func file_recommendation_proto_init() {
	if File_recommendation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommendation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelatedBlogReq); i {
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
		file_recommendation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelatedBlogRes); i {
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
		file_recommendation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRecommendedBlogsReq); i {
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
			RawDescriptor: file_recommendation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommendation_proto_goTypes,
		DependencyIndexes: file_recommendation_proto_depIdxs,
		MessageInfos:      file_recommendation_proto_msgTypes,
	}.Build()
	File_recommendation_proto = out.File
	file_recommendation_proto_rawDesc = nil
	file_recommendation_proto_goTypes = nil
	file_recommendation_proto_depIdxs = nil
}
