package handlers

import (
	"net/http"

	"github.com/dailoi280702/se121/api-gateway/internal/service/user"
	"github.com/go-chi/chi/v5"
)

func NewUserHanlder(userService user.UserServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/", handleMarkBlogAsReaded(userService))

	return c
}

func handleMarkBlogAsReaded(userService user.UserServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 	var res *
		// 	convertJsonApiToGrpc(w, r,
		// 		func() error {
		// 			var err error
		// 			res, err = userService.(context.Background(), &utils.Empty{})
		// 			return err
		// 		},
		// 		convertWithPostFunc(func() {
		// 			SendJson(w, res)
		// 		}))
	}
}
