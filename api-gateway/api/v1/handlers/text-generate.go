package handlers

import (
	"context"
	"net/http"

	tg "github.com/dailoi280702/se121/pkg/go/grpc/generated/text_generate"
	"github.com/go-chi/chi/v5"
)

func NewTextGenerateRoutes(tgc tg.TextGenerateServiceClient) chi.Router {
	r := chi.NewRouter()

	r.Get("/car-review", handleGetCarReview(tgc))
	r.Get("/blog-summarize", handleSummarizeBlog(tgc))

	return r
}

func handleGetCarReview(tgc tg.TextGenerateServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req tg.GenerateReviewReq
		var res *tg.ResString
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = tgc.GenerateCarReview(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}

func handleSummarizeBlog(tgc tg.TextGenerateServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req tg.GenerateBlogSummarizationReq
		var res *tg.ResString
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = tgc.GenerateBlogSummarization(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}
