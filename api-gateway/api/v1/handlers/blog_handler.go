package handlers

import (
	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/go-chi/chi/v5"
)

func NewBlogRoutes(blogService blog.BlogServiceClient) chi.Router {
	c := chi.NewRouter()
	return c
}
