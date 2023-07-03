package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/go-chi/chi/v5"
)

func NewTagRoutes(blogService blog.BlogServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/", handleGetAllTags(blogService))

	return c
}

func handleGetAllTags(blogService blog.BlogServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res *blog.GetAllTagsRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = blogService.GetAllTag(context.Background(), &blog.Empty{})
				return err
			},
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
