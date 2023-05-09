package server

import (
	"context"

	"github.com/dailoi280702/se121/comment_service/pkg/comment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	comment.UnimplementedCommentServiceServer
}

func NewServer() *server {
	return &server{}
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

func (s *server) GetAllComment(*comment.Empty, comment.CommentService_GetAllCommentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllComment not implemented")
}
