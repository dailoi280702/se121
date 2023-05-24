// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: text_generate.proto

package text_generate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TextGenerateServiceClient is the client API for TextGenerateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextGenerateServiceClient interface {
	GenerateCarReview(ctx context.Context, in *GenerateReviewReq, opts ...grpc.CallOption) (*ResString, error)
	GenerateBlogSummarization(ctx context.Context, in *GenerateBlogSummarizationReq, opts ...grpc.CallOption) (*ResString, error)
}

type textGenerateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextGenerateServiceClient(cc grpc.ClientConnInterface) TextGenerateServiceClient {
	return &textGenerateServiceClient{cc}
}

func (c *textGenerateServiceClient) GenerateCarReview(ctx context.Context, in *GenerateReviewReq, opts ...grpc.CallOption) (*ResString, error) {
	out := new(ResString)
	err := c.cc.Invoke(ctx, "/text_generate.TextGenerateService/GenerateCarReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerateServiceClient) GenerateBlogSummarization(ctx context.Context, in *GenerateBlogSummarizationReq, opts ...grpc.CallOption) (*ResString, error) {
	out := new(ResString)
	err := c.cc.Invoke(ctx, "/text_generate.TextGenerateService/GenerateBlogSummarization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextGenerateServiceServer is the server API for TextGenerateService service.
// All implementations must embed UnimplementedTextGenerateServiceServer
// for forward compatibility
type TextGenerateServiceServer interface {
	GenerateCarReview(context.Context, *GenerateReviewReq) (*ResString, error)
	GenerateBlogSummarization(context.Context, *GenerateBlogSummarizationReq) (*ResString, error)
	mustEmbedUnimplementedTextGenerateServiceServer()
}

// UnimplementedTextGenerateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextGenerateServiceServer struct {
}

func (UnimplementedTextGenerateServiceServer) GenerateCarReview(context.Context, *GenerateReviewReq) (*ResString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCarReview not implemented")
}
func (UnimplementedTextGenerateServiceServer) GenerateBlogSummarization(context.Context, *GenerateBlogSummarizationReq) (*ResString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBlogSummarization not implemented")
}
func (UnimplementedTextGenerateServiceServer) mustEmbedUnimplementedTextGenerateServiceServer() {}

// UnsafeTextGenerateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextGenerateServiceServer will
// result in compilation errors.
type UnsafeTextGenerateServiceServer interface {
	mustEmbedUnimplementedTextGenerateServiceServer()
}

func RegisterTextGenerateServiceServer(s grpc.ServiceRegistrar, srv TextGenerateServiceServer) {
	s.RegisterService(&TextGenerateService_ServiceDesc, srv)
}

func _TextGenerateService_GenerateCarReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerateServiceServer).GenerateCarReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/text_generate.TextGenerateService/GenerateCarReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerateServiceServer).GenerateCarReview(ctx, req.(*GenerateReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerateService_GenerateBlogSummarization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateBlogSummarizationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerateServiceServer).GenerateBlogSummarization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/text_generate.TextGenerateService/GenerateBlogSummarization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerateServiceServer).GenerateBlogSummarization(ctx, req.(*GenerateBlogSummarizationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TextGenerateService_ServiceDesc is the grpc.ServiceDesc for TextGenerateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextGenerateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "text_generate.TextGenerateService",
	HandlerType: (*TextGenerateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCarReview",
			Handler:    _TextGenerateService_GenerateCarReview_Handler,
		},
		{
			MethodName: "GenerateBlogSummarization",
			Handler:    _TextGenerateService_GenerateBlogSummarization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "text_generate.proto",
}
