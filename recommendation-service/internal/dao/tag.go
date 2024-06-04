package dao

import (
	"context"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/cache"
	"github.com/dailoi280702/se121/recommendation-service/client/blogservice"
	"github.com/dailoi280702/se121/recommendation-service/client/userservice"
	"github.com/dailoi280702/se121/user-service/userpb"
)

const defaultRecentBlogsLimitSize = 25

type TagDao interface {
	GetTagsFromBlogId(ctx context.Context, blogId int32) ([]*blog.BlogTags, error)
	GetUserTagsFromRecentActivity(ctx context.Context, userId string, limit int32) ([]*blog.BlogTags, error)
	GetLatestTags(ctx context.Context) ([]*blog.BlogTags, error)
}

type TagRepository struct {
	blogService blog.BlogServiceClient
	userService user.UserServiceClient
	cache       *cache.DualCache
}
