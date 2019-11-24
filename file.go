package gocacher

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var File cacher

type fileCacher struct {
	path string
	ext  string
}

func init() {
	File = &fileCacher{
		path: "./runtime/cache",
		ext:  "",
	}
}

func (fc *fileCacher) Init(config map[string]interface{}) cacher {
	if path, ok := config["path"].(string); ok {
		fc.path = path
	}
	if ext, ok := config["ext"].(string); ok {
		fc.ext = ext
	}
	return fc
}

func (fc *fileCacher) Clone(config map[string]interface{}) cacher {
	newCacher := &fileCacher{
		path: fc.path,
		ext:  fc.ext,
	}
	return newCacher.Init(config)
}

func (*fileCacher) Clear() error {
	panic("implement me")
}

func (*fileCacher) Len() int {
	panic("implement me")
}

func (fc *fileCacher) Set(key string, value interface{}) error {
	return fc.SetExpire(key, value, 0)
}

func (fc *fileCacher) SetExpire(key string, value interface{}, exp time.Duration) error {
	if err := fc.fs(); err != nil {
		return err
	}
	ci := &cacherItem{
		key: key,
		val: value,
		exp: time.Now().Add(exp),
	}
	// exp 不大于 0 时，为永久缓存
	if exp <= 0 {
		ci.exp = farTime
	}

	//
	path, err := fc.filename(key)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, gobEncode(ci), os.ModePerm)
}

func (*fileCacher) Has(key string) bool {
	panic("implement me")
}

func (*fileCacher) Get(key string) (interface{}, error) {
	panic("implement me")
}

func (*fileCacher) Pull(key string) (interface{}, error) {
	panic("implement me")
}

func (*fileCacher) Remove(key string) bool {
	panic("implement me")
}

func (*fileCacher) Keys() []string {
	panic("implement me")
}

// 初始化文件目录
func (fc *fileCacher) fs() error {
	return os.MkdirAll(fc.path, os.ModePerm)
}

// 获取文件名
func (fc *fileCacher) filename(key string) (path string, err error) {
	path = strings.Trim(fc.path+"/"+key+"."+strings.Trim(fc.ext, "."), ".")
	err = os.MkdirAll(path, os.ModePerm)
	return
}
