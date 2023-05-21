package server

import (
	"context"
	"database/sql"

	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *sql.DB
	comment.UnimplementedCommentServiceServer
}

func NewServer(db *sql.DB) *server {
	return &server{db: db}
}

func (s *server) CreateComment(context.Context, *comment.CreateCommentReq) (*comment.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}

func (s *server) UpdateComment(context.Context, *comment.UpdateCommentReq) (*comment.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}

func (s *server) DeleteComment(context.Context, *comment.DeleteCommentReq) (*comment.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

func (s *server) GetComment(context.Context, *comment.GetCommentReq) (*comment.Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}

func (s *server) GetBlogComments(ctx context.Context, in *comment.GetBlogCommentsReq) (*comment.GetBlogCommentsRes, error) {
	return nil, nil
}
