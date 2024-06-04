package dao

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/client/blogservice"
)

type BlogDao interface {
	GetBlogFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error)
}

type BlogRepository struct {
	blogService blog.BlogServiceClient
}

func (r *BlogRepository) GetBlogFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error) {
	res, err := r.blogService.GetBlogsFromIds(ctx, &blog.BlogIds{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	return res.Blogs, nil
}

func NewBlogRepository() BlogDao {
	return &BlogRepository{
		blogService: blogservice.GetInstance(),
	}
}
