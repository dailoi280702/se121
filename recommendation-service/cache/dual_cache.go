package cache

import (
	"log"
	"sync"
	"time"
)

const (
	defaultMaxRetries = 10
	defaultBaseDelay  = 5 * time.Second
	defaultTimeLife   = 5 * time.Minute
)

var (
	once     sync.Once
	instance *DualCache
)

type DualCache struct {
	cache      Cache
	mu         sync.Mutex
	maxRetries int
	baseDelay  time.Duration
	isRetrying bool
}

func GetInstance() *DualCache {
	once.Do(func() {
		instance = &DualCache{
			cache:      redisClientInstance,
			maxRetries: defaultMaxRetries,
			baseDelay:  defaultBaseDelay,
		}
	})

	return instance
}

func (c *DualCache) SetCache(cache Cache) {
	c.cache = cache
}

func (c *DualCache) Set(key string, val any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.cache.Set(key, val, defaultTimeLife)
	if err != nil && !c.cache.IsAlive() {
		if c.isRetrying {
			return err
		} else {
			c.SetCache(GetMemCacheInstance())
			log.Println("Switched to MemCache")
			go c.retry()

			return c.cache.Set(key, val, defaultTimeLife)
		}
	}

	return err
}

func (c *DualCache) Get(key string, dest any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.cache.Get(key, dest)
	if err != nil && !c.cache.IsAlive() {
		if c.isRetrying {
			return err
		} else {
			c.SetCache(GetMemCacheInstance())
			log.Println("Switched to MemCache")
			go c.retry()

			return c.cache.Get(key, dest)
		}
	}

	return err
}

func (c *DualCache) sync() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	m, err := c.cache.GetAll("*")
	if err != nil {
		return err
	}

	redisClient := GetRedisInstance()

	for k, v := range m {
		if err := redisClient.Set(k, v, defaultTimeLife); err != nil {
			return err
		}
	}

	return nil
}

func (c *DualCache) retry() {
	c.isRetrying = true
	redisClient := GetRedisInstance()

	for i := 0; i < c.maxRetries; i++ {
		log.Println("attempt: ", i+1)
		if redisClient.IsAlive() {
			if err := c.sync(); err != nil {
				log.Printf("error caches synchronization %+v", err)
				break
			}

			log.Println("Switched to redis")
			_ = c.cache.Clear("*")
			break
		}

		time.Sleep(c.baseDelay)
	}

	c.isRetrying = false
	c.SetCache(redisClient)
}

// func calculateDelay(baseDelay time.Duration, backoff float64, attempt int) time.Duration {
// 	// Introduce some randomness to avoid synchronization issues
// 	randomFactor := rand.Float64()
// 	waitTime := time.Duration(baseDelay.Seconds() * math.Pow(backoff, float64(attempt)) * randomFactor)
// 	return waitTime
// }
