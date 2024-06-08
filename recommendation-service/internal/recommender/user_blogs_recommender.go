package recommender

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
)

type UserBlogsRecommender struct {
	DefaultBlogRecommender
}

func (r *UserBlogsRecommender) GetInteractedBlogTags(ctx context.Context, id string, limit int32) ([]*blog.BlogTags, error) {
	return r.tagRepo.GetUserTagsFromRecentActivity(ctx, id, limit)
}

func NewUserBlogsRecommender(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) BlogRecommender {
	defaultRecommender := &DefaultBlogRecommender{
		blogRepo: blogRepo,
		tagRepo:  tagRepo,
	}

	recommender := &UserBlogsRecommender{
		DefaultBlogRecommender: *defaultRecommender,
	}

	defaultRecommender.BlogRecommender = recommender

	return defaultRecommender
}
