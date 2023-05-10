package server

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateTag(context.Context, *blog.CreateTagReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}

func (s *server) UpdateTag(context.Context, *blog.UpdateTagReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}

func (s *server) DeleteTag(context.Context, *blog.DeleteTagReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}

func (s *server) GetTag(context.Context, *blog.GetTagReq) (*blog.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}

func (s *server) GetAllTag(*blog.Empty, blog.BlogService_GetAllTagServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllTag not implemented")
}
