package adapter

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheCacher struct {
	client *memcache.Client
}

func NewMemcacheCacher(server ...string) *MemcacheCacher {
	return &MemcacheCacher{
		client: memcache.New(server...),
	}
}

// Set
func (mc *MemcacheCacher) Set(key string, val interface{}, timeout time.Duration) error {
	item := &CacheItem{
		Key:      key,
		Data:     val,
		CreateAt: time.Now(),
		Exp:      timeout * time.Second,
	}
	return mc.client.Set(&memcache.Item{
		Key:        key,
		Value:      item.GobEncode(),
		Expiration: int32(timeout),
	})
}

// Get
func (mc *MemcacheCacher) Get(key string) interface{} {
	item, err := mc.client.Get(key)
	if err != nil {
		return nil
	}
	data := item.Value
	cache := &CacheItem{}
	cache.GobDecode(data)
	return cache.Data
}

// Remove
func (mc *MemcacheCacher) Remove(key string) error {
	return mc.client.Delete(key)
}

// Pull
func (mc *MemcacheCacher) Pull(key string) interface{} {
	defer mc.Remove(key)
	return mc.Get(key)
}

// Clear
func (mc *MemcacheCacher) Clear() error {
	return mc.client.DeleteAll()
}
