package server

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateTag(context.Context, *blog.CreateTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}

func (s *server) UpdateTag(context.Context, *blog.UpdateTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}

func (s *server) DeleteTag(context.Context, *blog.DeleteTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}

func (s *server) GetTag(context.Context, *blog.GetTagReq) (*blog.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}

func (s *server) GetAllTag(context.Context, *utils.Empty) (*blog.GetAllTagsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTag not implemented")
}
