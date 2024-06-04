package cache

import "time"

type Cache interface {
	Set(key string, val any, expiration time.Duration) error
	Get(key string, dest any) error
	Del(key string) error
	GetAll(pattern string) (map[string]any, error)
	Clear(pattern string) error
	IsAlive() bool
}
