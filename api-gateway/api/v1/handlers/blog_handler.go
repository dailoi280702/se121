package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/go-chi/chi/v5"
)

func NewBlogRoutes(blogService blog.BlogServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/", handleGetBlog(blogService))
	c.Put("/", handleUpdateBlog(blogService))
	c.Post("/", handleCreateBlog(blogService))
	c.Delete("/", handleDeleteBlog(blogService))
	c.Get("/Search", handleSearchBlog(blogService))

	return c
}

func handleGetBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.GetBlogReq
		var res *blog.Blog
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.GetBlog(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleCreateBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.CreateBlogReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := blogService.CreateBlog(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleUpdateBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.UpdateBlogReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := blogService.UpdateBlog(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleDeleteBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.DeleteBlogReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := blogService.DeleteBlog(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleSearchBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.SearchReq
		var res *blog.SearchBlogsRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.SearchForBlogs(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
