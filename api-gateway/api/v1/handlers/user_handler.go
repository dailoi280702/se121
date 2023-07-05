package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/user-service/userpb"
	"github.com/go-chi/chi/v5"
)

func NewUserHanlder(userService user.UserServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Post("/readed-blog", handleMarkBlogAsReaded(userService))

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
