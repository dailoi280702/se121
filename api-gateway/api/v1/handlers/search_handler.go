package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/dailoi280702/se121/search-service/pkg/search"
	"github.com/go-chi/chi/v5"
)

func NewSearchRoutes(searchService search.SearchServiceClient) chi.Router {
	r := chi.NewRouter()

	r.Get("/", handleSearch(searchService))

	return r
}

func handleSearch(searchService search.SearchServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req utils.SearchReq
		var res *search.SearchRes
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = searchService.Search(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}
