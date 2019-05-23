package adapter

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Adapter interface {
	Set(key string, val interface{}, timeout time.Duration) error
	Get(key string) interface{}
	Remove(key string) error
	Pull(key string) interface{}
	Clear() error
	// All(keys []string) []interface{}
	// Tag(tag string, key string, val interface{}, timeout time.Duration) *Adapter
	// Inc(key string) error
	// Dec(key string) error
}

type CacheItem struct {
	Key      string
	Data     interface{}
	CreateAt time.Time
	Exp      time.Duration
}

func NewCacheItem(key string, data interface{}, exp time.Duration) *CacheItem {
	return &CacheItem{
		Key:      key,
		Data:     data,
		CreateAt: time.Now(),
		Exp:      exp,
	}
}

type CacherType string

const (
	TypeFile     CacherType = "file"
	TypeMemory   CacherType = "memory"
	TypeRedis    CacherType = "redis"
	TypeMemcache CacherType = "memcache"
)

// 验证过期，过期返回true
func (item *CacheItem) isExpired() bool {
	return item.Exp != 0 && item.CreateAt.Add(item.Exp).Before(time.Now())
}

// Go Gob 序列化
func (data *CacheItem) GobEncode() []byte {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

// Go Gob 反序列化
func (to *CacheItem) GobDecode(data []byte) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(&to)
}
