package cache

import (
	"encoding/json"
	"errors"
	"regexp"
	"sync"
	"time"
)

type item struct {
	Value  any
	Expiry time.Time
}

func (i item) isExpired() bool {
	return time.Now().After(i.Expiry)
}

type MemCache struct {
	items map[string]item
	mu    sync.Mutex
}

var (
	memCacheInstance *MemCache
	initMemCache     sync.Once
)

func GetMemCacheInstance() Cache {
	initMemCache.Do(func() {
		memCacheInstance = &MemCache{
			items: map[string]item{},
		}
	})

	return memCacheInstance
}

func (m *MemCache) Set(key string, val any, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[key] = item{
		Value:  val,
		Expiry: time.Now().Add(expiration),
	}

	return nil
}

func (m *MemCache) Get(key string, dest any) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	item, found := m.items[key]
	if !found {
		return errors.New("key " + key + " not found")
	}

	if item.isExpired() {
		delete(m.items, key)
		return errors.New("key " + key + " is expired")
	}

	tmp, _ := json.Marshal(item.Value)
	return json.Unmarshal(tmp, dest)
}

func (m *MemCache) Del(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.items, key)

	return nil
}

func (m *MemCache) GetAll(pattern string) (map[string]any, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	res := map[string]any{}
	for k, v := range m.items {
		if v.isExpired() {
			delete(m.items, k)
		} else {
			res[k] = v.Value
		}
	}

	return res, nil
}

func (m *MemCache) Clear(pattern string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k := range m.items {
		if matched, err := regexp.MatchString(pattern, k); err == nil && matched {
			delete(m.items, k)
		}
	}

	return nil
}

func (m *MemCache) IsAlive() bool {
	return true
}
