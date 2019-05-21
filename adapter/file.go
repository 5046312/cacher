package adapter

import "time"

const (
	FILE_CACHER_PATH string = "./runtime/cache/"
	FILE_CACHER_EXT  string = ".cache"
)

type FileCacher struct {
	Path string
	Ext  string
}

type CacheFile struct {
	Key  string
	Data interface{}
	Exp  time.Duration
}

//
func NewFileCacher(path string, ext string) *FileCacher {
	fc := DefaultFileCacher()
	fc.setPath(path)
	fc.setExt(ext)
	return fc
}

// 默认配置
func DefaultFileCacher() *FileCacher {
	return &FileCacher{FILE_CACHER_PATH, FILE_CACHER_EXT}
}

// TODO
func (fc *FileCacher) setPath(path string) {
	fc.Path = path
}

// TODO
func (fc *FileCacher) setExt(ext string) {
	fc.Ext = ext
}

func (fc *FileCacher) Get(key string) interface{} {
	return nil
}
func (fc *FileCacher) All(keys []string) []interface{} {
	return nil
}
func (fc *FileCacher) Set(key string, val interface{}, timeout time.Duration) error {
	return nil
}
func (fc *FileCacher) Inc(key string) error {
	return nil
}
func (fc *FileCacher) Dec(key string) error {
	return nil
}
func (fc *FileCacher) Remove(key string) error {
	return nil
}
func (fc *FileCacher) Pull(key string) interface{} {
	return nil
}
func (fc *FileCacher) IsExist(key string) bool {
	return false
}
func (fc *FileCacher) Clear() error {
	return nil
}
