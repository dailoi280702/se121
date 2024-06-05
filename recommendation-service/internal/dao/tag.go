package dao

import (
	"context"
	"log"

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

func (r *TagRepository) GetTagsFromBlogId(ctx context.Context, blogId int32) ([]*blog.BlogTags, error) {
	res, err := r.blogService.GetTagsFromBlogIds(ctx, &blog.BlogIds{
		Ids: []int32{blogId},
	})
	if err != nil {
		return nil, err
	}

	return res.BlogTags, nil
}

func (r *TagRepository) GetUserTagsFromRecentActivity(ctx context.Context, userId string, limit int32) ([]*blog.BlogTags, error) {
	blogsRes, err := r.userService.GetRecentlyReadedBlogsIds(ctx, &user.GetRecentlyReadedBlogsIdsReq{
		UserId: userId,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	tagsRes, err := r.blogService.GetTagsFromBlogIds(ctx, &blog.BlogIds{Ids: blogsRes.BlogIds})
	if err != nil {
		return nil, err
	}

	return tagsRes.BlogTags, nil
}

func (r *TagRepository) GetLatestTags(ctx context.Context) ([]*blog.BlogTags, error) {
	res := &blog.GetLatestBlogTagsRes{}
	key := "blogs:latest"

	if r.cache != nil {
		if err := r.cache.Get(key, res); err == nil {
			log.Println("cache hit, key: ", key)
			return res.GetBlogTags(), nil
		}

		log.Println("cache miss, key: ", key)
	}

	res, err := r.blogService.GetLatestBlogTags(ctx, &blog.GetLatestBlogTagsReq{
		GetNumberOfBlogs: defaultRecentBlogsLimitSize,
	})
	if err != nil {
		return nil, err
	}

	if r.cache != nil {
		go func() {
			if err := r.cache.Set(key, res); err != nil {
				log.Printf("Error saving to cache: %+v", err)
			}
		}()
	}

	return res.BlogTags, nil
}

func NewTagRepository() TagDao {
	return &TagRepository{
		blogService: blogservice.GetInstance(),
		userService: userservice.GetInstance(),
		cache:       cache.GetInstance(),
	}
}

// func blogTagsToTags(blogTags []*blog.BlogTags) []*blog.Tag {
// 	tags := []*blog.Tag{}
// 	mTag := map[int32]*blog.Tag{}
//
// 	for i := range blogTags {
// 		if blogTags[i] != nil {
// 			currentTags := blogTags[i].GetTags()
// 			for j := range currentTags {
// 				if _, ok := mTag[currentTags[j].Id]; !ok {
// 					tags = append(tags, currentTags[j])
// 				}
// 			}
// 		}
// 	}
//
// 	return tags
// }
