package repo

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/client/blogservice"
)

type BlogRepository interface {
	GetBlogFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error)
}

type blogRepository struct {
	blogService blog.BlogServiceClient
}

func (r *blogRepository) GetBlogFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error) {
	res, err := r.blogService.GetBlogsFromIds(ctx, &blog.BlogIds{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	return res.Blogs, nil
}

func NewBlogRepository() BlogRepository {
	return &blogRepository{
		blogService: blogservice.GetInstance(),
	}
}
