package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
	"github.com/dailoi280702/se121/user-service/userpb"
	"github.com/go-chi/chi/v5"
)

func NewUserHanlder(userService user.UserServiceClient, recommendedService recommendation.RecommendationServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Post("/readed-blog", handleMarkBlogAsReaded(userService))
	c.Get("/{id}/recommended-blogs", handleGetUserRecommendedBlogs(recommendedService))

	return c
}

func handleMarkBlogAsReaded(userService user.UserServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req user.MarkBlogAsReadedReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := userService.MarkBlogAsReaded(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleGetUserRecommendedBlogs(recommendedService recommendation.RecommendationServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		req := recommendation.GetUserRecommendedBlogsReq{UserId: id}
		var res *blog.Blogs
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = recommendedService.GetUserRecommendedBlogs(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}
