package gocacher

import (
	"io/ioutil"
	"log"
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
		ext:  "cache",
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

func (fc *fileCacher) Clear() error {
	return os.RemoveAll(fc.path)
}

func (fc *fileCacher) Len() int {
	files, _ := ioutil.ReadDir(fc.path)
	return len(files)
}

func (fc *fileCacher) Set(key string, value interface{}) error {
	return fc.SetExpire(key, value, 0)
}

func (fc *fileCacher) SetExpire(key string, value interface{}, exp time.Duration) error {
	if err := fc.fs(); err != nil {
		return err
	}
	ci := &cacherItem{
		Key: key,
		Val: value,
		Exp: time.Now().Add(exp),
	}
	// exp 不大于 0 时，为永久缓存
	if exp <= 0 {
		ci.Exp = farTime
	}

	//
	filename := fc.filename(key)
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	encode := gobEncode(ci)
	file.WriteString(encode)
	return nil
}

func (fc *fileCacher) Has(key string) bool {
	_, err := fc.Get(key)
	return err == nil
}

func (fc *fileCacher) Get(key string) (interface{}, error) {
	filename := fc.filename(key)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	ci, err := gobDecode(content)
	if ci.expired() {
		fc.Remove(key)
		return nil, KeyExpireError
	}
	return ci.Val, nil
}

func (fc *fileCacher) Pull(key string) (interface{}, error) {
	val, err := fc.Get(key)
	if err != nil {
		return val, err
	}
	fc.Remove(key)
	return val, nil
}

func (fc *fileCacher) Remove(key string) bool {
	os.Remove(fc.filename(key))
	return true
}

func (fc *fileCacher) Keys() []string {
	keys := []string{}
	files, _ := ioutil.ReadDir(fc.path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		file, err := os.Open(strings.Trim(fc.path, "/") + "/" + file.Name())
		defer file.Close()
		if err != nil {
			log.Fatalln("读取文件失败")
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln("读取文件内容失败")
		}
		ci, err := gobDecode(content)
		if ci.expired() {
			fc.Remove(ci.Key)
		} else {
			keys = append(keys, ci.Key)
		}
	}
	return keys
}

// 初始化文件目录
func (fc *fileCacher) fs() error {
	return os.MkdirAll(fc.path, os.ModePerm)
}

// 获取文件名
func (fc *fileCacher) filename(key string) string {
	return fc.path + "/" + md5Encode(key) + "." + strings.Trim(fc.ext, ".")
}
