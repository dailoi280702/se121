package proxyserver

import (
	"context"
	"fmt"
	"log"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/cache"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
)

var _ recommendation.RecommendationServiceServer = &ProxyServer{}

type ProxyServer struct {
	recommendation.UnimplementedRecommendationServiceServer
	cache  *cache.DualCache
	server recommendation.RecommendationServiceServer
}

func NewProxyServer(server recommendation.RecommendationServiceServer, cache *cache.DualCache) recommendation.RecommendationServiceServer {
	return &ProxyServer{
		server: server,
		cache:  cache,
	}
}

func (s *ProxyServer) GetRelatedBlog(ctx context.Context, req *recommendation.GetRelatedBlogReq) (*recommendation.GetRelatedBlogRes, error) {
	res := &recommendation.GetRelatedBlogRes{}
	key := fmt.Sprintf("recommendation:blog:%d:size:%d", req.BlogId, req.NumberOfBlog)

	if s.cache != nil {
		if err := s.cache.Get(key, res); err == nil {
			log.Println("cache hit, key: ", key)
			return res, nil
		}

		log.Println("cache miss, key: ", key)
	}

	res, err := s.server.GetRelatedBlog(ctx, req)
	if err != nil {
		log.Printf("Error get related blog %+v", err)

		return nil, fmt.Errorf("Blog server error: %v", err)
	}

	if s.cache != nil {
		go func() {
			if err := s.cache.Set(key, res); err != nil {
				log.Printf("Error saving to cache: %+v", err)
			}
		}()
	}

	return res, nil
}

func (s *ProxyServer) GetUserRecommendedBlogs(ctx context.Context, req *recommendation.GetUserRecommendedBlogsReq) (*blog.Blogs, error) {
	res := &blog.Blogs{}
	key := fmt.Sprintf("recommendation:user:%s:limit:%d", req.UserId, req.Limit)

	if s.cache != nil {
		if err := s.cache.Get(key, res); err == nil {
			log.Println("cache hit, key: ", key)
			return res, nil
		}

		log.Println("cache miss, key: ", key)
	}
	res, err := s.server.GetUserRecommendedBlogs(ctx, req)
	if err != nil {
		log.Printf("Error get user recommendation %+v", err)

		return nil, fmt.Errorf("Blog server error: %v", err)
	}

	if s.cache != nil {
		go func() {
			if err := s.cache.Set(key, res); err != nil {
				log.Printf("Error saving to cache: %+v", err)
			}
		}()
	}

	return res, nil
}
