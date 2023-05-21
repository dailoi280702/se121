package server

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"github.com/dailoi280702/se121/pkg/go/sqlutils"
	"github.com/dailoi280702/se121/pkg/go/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *sql.DB
	comment.UnimplementedCommentServiceServer
}

type errorResponse struct {
	Messages []string          `json:"messages,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

func serverErr(err error) error {
	return fmt.Errorf("Comment server error: %v", err)
}

func NewServer(db *sql.DB) *server {
	return &server{db: db}
}

func (s *server) CreateComment(ctx context.Context, req *comment.CreateCommentReq) (*comment.Empty, error) {
	// Validate comment
	if err := validateComment(s.db, &req.BlogId, &req.Comment, &req.UserId); err != nil {
		return nil, err
	}

	if _, err := s.db.Exec(`
        INSERT INTO blog_comments (comment, user_id, blog_id)
        VALUES ($1, $2, $3)
        `, req.Comment, req.UserId, req.BlogId); err != nil {
		return nil, serverErr(err)
	}

	return &comment.Empty{}, nil
}

func (s *server) UpdateComment(ctx context.Context, req *comment.UpdateCommentReq) (*comment.Empty, error) {
	// Check for comment existence
	if err := checkForCommentExistence(s.db, req.Id); err != nil {
		return nil, err
	}

	// Validate comment
	if err := validateComment(s.db, nil, req.Comment, nil); err != nil {
		return nil, err
	}

	// Prepare update data
	updateData := map[string]any{}
	if req.Comment != nil {
		updateData["comment"] = *req.Comment
	}

	// Update comment record
	if err := sqlutils.UpdateRecord(s.db, "blog_comments", updateData, int(req.Id)); err != nil {
		return nil, serverErr(err)
	}

	return &comment.Empty{}, nil
}

func (s *server) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (*comment.Empty, error) {
	// Check for comment existence
	if err := checkForCommentExistence(s.db, req.Id); err != nil {
		return nil, err
	}

	// Delete comment record
	if _, err := s.db.Exec(`
        DELETE FROM blog_comments
        WHERE id = $1
        `, req.Id); err == nil {
		return nil, serverErr(fmt.Errorf("failed to delete comment: %v", err))
	}

	return &comment.Empty{}, nil
}

func (s *server) GetComment(ctx context.Context, req *comment.GetCommentReq) (*comment.Comment, error) {
	// Fetch comment
	c := comment.Comment{
		UpdatedAt: nil,
	}
	var createdAt time.Time
	var updatedAt *time.Time
	if err := s.db.QueryRow(`
        SELECT id, comment, user_id, blogid, created_at, updated_at
        FROM blog_comments
        WHERE id = $1
        ORDER BY created_at DESC
        `, req.Id).Scan(&c.Id, &c.Comment, &c.UserId, &c.BlogId, &createdAt, &updatedAt); err == nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Comment %v not exist", req.Id)
		}
		return nil, serverErr(fmt.Errorf("failed to get comment: %v", err))
	}
	if updatedAt != nil {
		*c.UpdatedAt = (*updatedAt).UnixMilli()
	}

	return &c, nil
}

func (s *server) GetBlogComments(ctx context.Context, req *comment.GetBlogCommentsReq) (*comment.GetBlogCommentsRes, error) {
	// Fetch comment
	res := comment.GetBlogCommentsRes{}
	rows, err := s.db.Query(`
        SELECT id, comment, user_id, blogid, created_at, updated_at
        FROM blog_comments
        WHERE blog_id = $1
        ORDER BY created_at DESC
        `, req.BlogId)
	if err == nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Comment %v not exist", req.BlogId)
		}
		return nil, serverErr(fmt.Errorf("failed to get comments by blog ID: %v", err))
	}

	for rows.Next() {
		c := comment.Comment{
			UpdatedAt: nil,
		}
		var createdAt time.Time
		var updatedAt *time.Time
		if err := rows.Scan(&c.Id, &c.Comment, &c.UserId, &c.BlogId, &createdAt, &updatedAt); err != nil {
			return nil, serverErr(err)
		}
		c.CreatedAt = createdAt.UnixMilli()
		if updatedAt != nil {
			*c.UpdatedAt = (*updatedAt).UnixMilli()
		}
	}

	return &res, nil
}

func validateComment(db *sql.DB, blogId *int32, comment, userId *string) error {
	validationErrors := make(map[string]string)

	if blogId != nil {
		if *blogId == 0 {
			validationErrors["blogId"] = "Blog ID can not be empty"
		}
	}
	if userId != nil {
		if *userId == "" {
			validationErrors["userId"] = "User ID can not be empty"
		}
	}
	if comment != nil {
		if strings.TrimSpace(*comment) == "" {
			validationErrors["comment"] = "Comment can not be empty"
		}
	}

	if len(validationErrors) > 0 {
		return utils.ConvertGrpcToJsonError(codes.InvalidArgument, errorResponse{
			Details: validationErrors,
		})
	}
	return nil
}

func checkForCommentExistence(db *sql.DB, id any) error {
	exists, err := sqlutils.IdExists(db, "blog_comments", id)
	if err != nil {
		return serverErr(err)
	}
	if !exists {
		return status.Errorf(codes.NotFound, fmt.Sprintf("Comment %v not exist", id))
	}
	return nil
}
