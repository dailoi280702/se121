package recommender

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
)

type UserBlogsDataSource struct {
	blogRepo repo.BlogRepository
	tagRepo  repo.TagRepository
}

func (r *UserBlogsDataSource) GetRelatedBlogTags(ctx context.Context, id string, limit int32) ([]*blog.BlogTags, error) {
	return r.tagRepo.GetUserTagsFromRecentActivity(ctx, id, limit)
}

func (r *UserBlogsDataSource) GetLatestBlogTags(ctx context.Context) ([]*blog.BlogTags, error) {
	return r.tagRepo.GetLatestTags(ctx)
}

func (r *UserBlogsDataSource) GetBlogsFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error) {
	return r.blogRepo.GetBlogFromIds(ctx, ids)
}

func NewUserBlogsRecommender(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) RecommenderDataSource {
	return &UserBlogsDataSource{
		blogRepo: blogRepo,
		tagRepo:  tagRepo,
	}
}
