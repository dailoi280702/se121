package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"github.com/go-chi/chi/v5"
)

func NewCommentRoutes(commentService comment.CommentServiceClient) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handleCreateComment(commentService))
	r.Put("/", handleUpdateComment(commentService))
	r.Get("/", handleGetCommentById(commentService))
	r.Get("/blog/{blogId}", handleGetCommentByBlogId(commentService))
	r.Delete("/", handleDeleteCommentById(commentService))
	return r
}

func handleUpdateComment(commentService comment.CommentServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req comment.UpdateCommentReq
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				_, err = commentService.UpdateComment(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleDeleteCommentById(commentService comment.CommentServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req comment.DeleteCommentReq
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				_, err = commentService.DeleteComment(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleGetCommentById(commentService comment.CommentServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req comment.GetCommentReq
		var res *comment.Comment
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = commentService.GetComment(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleGetCommentByBlogId(commentService comment.CommentServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blogId, err := strconv.Atoi(chi.URLParam(r, "blogId"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := comment.GetBlogCommentsReq{BlogId: int32(blogId)}
		var res *comment.GetBlogCommentsRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = commentService.GetBlogComments(context.Background(), &req)
				return err
			},
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleCreateComment(commentService comment.CommentServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req comment.CreateCommentReq
		convertJsonApiToGrpc(w, r, func() error {
			var err error
			_, err = commentService.CreateComment(context.Background(), &req)
			return err
		}, convertWithJsonReqData(&req))
	}
}
