package adapter

import "time"

type Adapter interface {
	Get(key string) interface{}
	All(keys []string) []interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	// Tag(tag string, key string, val interface{}, timeout time.Duration) *Adapter
	Inc(key string) error
	Dec(key string) error
	Remove(key string) error
	Pull(key string) interface{}
	IsExist(key string) bool
	Clear() error
}

type CacherType string

const (
	TypeFile     CacherType = "file"
	TypeMemory   CacherType = "memory"
	TypeRedis    CacherType = "redis"
	TypeMemcache CacherType = "memcache"
)
