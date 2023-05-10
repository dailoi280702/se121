package server

import "github.com/dailoi280702/se121/blog-service/pkg/blog"

type server struct {
	blog.UnimplementedBlogServiceServer
}

func NewServer() *server {
	return &server{}
}
