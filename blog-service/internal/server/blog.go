package server

import (
	"context"

	"github.com/dailoi280702/se121/blog_service/pkg/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateBlog(context.Context, *blog.CreateBlogReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}

func (s *server) GetBlog(context.Context, *blog.GetBlogReq) (*blog.Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}

func (s *server) UpdateBlog(context.Context, *blog.UpdateBlogReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}

func (s *server) DeleteBlog(context.Context, *blog.DeleteBlogReq) (*blog.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}

func (s *server) SearchForBlogs(*blog.SearchForBlogsReq, blog.BlogService_SearchForBlogsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchForBlogs not implemented")
}

func (s *server) GetNumberOfBlogs(context.Context, *blog.Empty) (*blog.GetNumberOfBlogsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfBlogs not implemented")
}
