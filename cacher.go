package gocacher

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"time"
)

type cacher interface {
	Init(config map[string]interface{}) cacher
	Clone(config map[string]interface{}) cacher
	Set(key string, value interface{}) error
	SetExpire(key string, value interface{}, exp time.Duration) error
	Has(key string) bool
	Get(key string) (interface{}, error)
	Pull(key string) (interface{}, error)
	Remove(key string) bool
	Clear() error
	Keys() []string
	Len() int
}

type cacherItem struct {
	Key string
	Val interface{}
	Exp time.Time
}

// 该数据是否过期
func (ci *cacherItem) expired() bool {
	return ci.Exp.Before(time.Now())
}

// Gob Encode
func gobEncode(ci *cacherItem) (string, error) {
	gob.Register(ci.Val)
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(ci)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Gob Decode
func gobDecode(data []byte) (ci *cacherItem, err error) {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&ci)
	return
}

// Md5
func md5Encode(ori string) string {
	h := md5.New()
	h.Write([]byte(ori))
	return hex.EncodeToString(h.Sum(nil))
}

var (
	gcTime  = time.Hour
	farTime = time.Date(3018, 11, 23, 22, 44, 0, 0, time.Local)
)

var (
	KeyNotExistError = errors.New("key not exist")
	KeyExpireError   = errors.New("key expired")
)
