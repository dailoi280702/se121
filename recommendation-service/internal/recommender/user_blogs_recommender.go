package recommender

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/dao"
)

type UserBlogsRecommender struct {
	DefaultBlogRecommender
}

func (r *UserBlogsRecommender) GetInteractedBlogTags(ctx context.Context, id string, limit int32) ([]*blog.BlogTags, error) {
	return r.tagDao.GetUserTagsFromRecentActivity(ctx, id, limit)
}

func NewUserBlogsRecommender(blogDao dao.BlogDao, tagDao dao.TagDao) BlogRecommender {
	defaultRecommender := &DefaultBlogRecommender{
		blogDao: blogDao,
		tagDao:  tagDao,
	}

	recommender := &UserBlogsRecommender{
		DefaultBlogRecommender: *defaultRecommender,
	}

	defaultRecommender.BlogRecommender = recommender

	return defaultRecommender
}
