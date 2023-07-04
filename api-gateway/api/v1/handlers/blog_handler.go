package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/go-chi/chi/v5"
)

func NewBlogRoutes(blogService blog.BlogServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/{id}", handleGetBlog(blogService))
	c.Put("/", handleUpdateBlog(blogService))
	c.Post("/", handleCreateBlog(blogService))
	c.Delete("/", handleDeleteBlog(blogService))
	c.Get("/search", handleSearchBlog(blogService))

	return c
}

func handleGetBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := blog.GetBlogReq{Id: int32(id)}
		var res *blog.Blog
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.GetBlog(context.Background(), &req)
				return err
			},
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleCreateBlog(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req blog.CreateBlogReq
		var res *blog.CreateBlogRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.CreateBlog(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
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
		var req utils.SearchReq
		var res *blog.SearchBlogsRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.SearchForBlogs(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
