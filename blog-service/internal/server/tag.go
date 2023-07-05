package server

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateTag(context.Context, *blog.CreateTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}

func (s *server) UpdateTag(context.Context, *blog.UpdateTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}

func (s *server) DeleteTag(context.Context, *blog.DeleteTagReq) (*utils.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}

func (s *server) GetTag(context.Context, *blog.GetTagReq) (*blog.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}

func (s *server) GetAllTag(context.Context, *utils.Empty) (*blog.GetAllTagsRes, error) {
	tags := []*blog.Tag{}

	rows, err := s.db.Query(`
        SELECT id, name, description from tags
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return &blog.GetAllTagsRes{Tags: tags}, nil
		}
		return nil, serverError(fmt.Errorf("error while getting tags from db: %v", err))
	}

	for rows.Next() {
		var tag blog.Tag
		err := rows.Scan(&tag.Id, &tag.Name, &tag.Description)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	defer rows.Close()

	return &blog.GetAllTagsRes{Tags: tags}, nil
}
