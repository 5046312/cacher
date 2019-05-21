package adapter

import (
	"bytes"
	"encoding/gob"
	"time"
)

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

// Go Gob 序列化
func GobEncode(data interface{}) []byte {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

// Go Gob 反序列化
func GobDecode(data []byte, to interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(&to)
}
