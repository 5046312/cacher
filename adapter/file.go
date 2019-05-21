package adapter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	FILE_CACHER_PATH string = "./runtime/cache/"
	FILE_CACHER_EXT  string = ""
)

type FileCacher struct {
	Path string
	Ext  string
}

//
func NewFileCacher(path string) *FileCacher {
	fc := DefaultFileCacher()
	fc.setPath(path)
	return fc
}

// 默认配置
func DefaultFileCacher() *FileCacher {
	fc := &FileCacher{FILE_CACHER_PATH, FILE_CACHER_EXT}
	fc.setPath(FILE_CACHER_PATH)
	return fc
}

// Set File Cache Path
func (fc *FileCacher) setPath(path string) {
	fc.Path = path
	// Create Fold
	os.MkdirAll(fc.Path, os.ModePerm)
}

// Set File Cache Ext
func (fc *FileCacher) setExt(ext string) {
	fc.Ext = "." + strings.Trim(ext, ".")
}

// Get Cache File Name
func (fc *FileCacher) getCacheFileName(key string) string {
	m := md5.New()
	io.WriteString(m, key)
	md5Key := hex.EncodeToString(m.Sum(nil))
	cachePath := filepath.Join(fc.Path, md5Key[10:12], md5Key[20:22])
	os.MkdirAll(cachePath, os.ModePerm)
	return filepath.Join(cachePath, md5Key+fc.Ext)
}

//
func (fc *FileCacher) Get(key string) interface{} {
	filename := fc.getCacheFileName(key)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	cache := &CacheItem{}
	err = GobDecode(data, cache)
	if err != nil || cache.Exp.Before(time.Now()) {
		fc.Remove(key)
		return nil
	}
	return cache.Data
}

//
func (fc *FileCacher) All(keys []string) []interface{} {
	return nil
}

// Set
func (fc *FileCacher) Set(key string, val interface{}, timeout time.Duration) error {
	filename := fc.getCacheFileName(key)
	fmt.Println(filename)
	cache := &CacheItem{
		Key:  key,
		Data: val,
		Exp:  time.Now().Add(timeout * time.Second),
	}
	return ioutil.WriteFile(filename, GobEncode(cache), os.ModePerm)
}

//
func (fc *FileCacher) Inc(key string) error {
	return nil
}

//
func (fc *FileCacher) Dec(key string) error {
	return nil
}

//
func (fc *FileCacher) Remove(key string) error {
	filename := fc.getCacheFileName(key)
	return os.Remove(filename)
}

//
func (fc *FileCacher) Pull(key string) interface{} {
	cache := fc.Get(key)
	fc.Remove(key)
	return cache
}

//
func (fc *FileCacher) Clear() error {
	return os.RemoveAll(fc.Path)
}
