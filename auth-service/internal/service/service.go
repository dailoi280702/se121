package service

import (
	"github.com/redis/go-redis/v9"
)

type Service struct {
	rdb *redis.Client
}
