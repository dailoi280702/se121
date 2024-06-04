package cache

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisAddr = flag.String("redisAddr", "redis:6379", "the address to connect to redis")

var (
	redisClientInstance *RedisClient
	initRedisClient     sync.Once
)

type RedisClient struct {
	C   *redis.Client
	ctx context.Context
}

func GetRedisInstance() Cache {
	initRedisClient.Do(func() {
		c := redis.NewClient(&redis.Options{
			Addr:     *redisAddr,
			Password: "",
			DB:       0,
		})
		redisClientInstance = &RedisClient{C: c, ctx: context.Background()}

		if err := redisClientInstance.C.Ping(context.Background()).Err(); err != nil {
			log.Println("Redis server is down")
		}
	})

	return redisClientInstance
}

func (r *RedisClient) Set(key string, value any, expiration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.C.Set(r.ctx, key, p, expiration).Err()
}

func (r *RedisClient) Get(key string, dest any) error {
	p, err := r.C.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(p), dest)
}

func (r *RedisClient) Del(key string) error {
	return r.C.Del(r.ctx, key).Err()
}

func (r *RedisClient) GetAll(pattern string) (map[string]any, error) {
	return map[string]any{}, nil
}

func (r *RedisClient) Clear(pattern string) error {
	keys, err := r.C.Keys(context.Background(), pattern).Result()
	if err != nil {
		log.Println("Error deleting key:", err)
	}

	for _, key := range keys {
		if err := r.C.Del(context.Background(), key).Err(); err != nil {
			log.Println("Error deleting key:", key, err)
		}
	}

	return nil
}

func (r *RedisClient) IsAlive() bool {
	err := r.C.Ping(r.ctx).Err()
	if err != nil {
		log.Println("Redis server is down")
	}

	return err == nil
}
