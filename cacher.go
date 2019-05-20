package cacher

import "github.com/5046312/cacher/adapter"

type Cacher struct {
	adapter map[adapter.CacherType]adapter.Adapter
}

var cacher *Cacher

func getCacher() *Cacher {
	if cacher == nil {
		cacher = &Cacher{}
	}
	return cacher
}

func FileConfig() *adapter.CacherFile {
	return adapter.DefaultCacherFileConfig()
}

func SetFile(fc *adapter.CacherFile) adapter.Adapter {
	// 判断是否已经加载过对应adapter
	if _, ok := getCacher().adapter[adapter.TypeFile]; ok {
		panic("Cacher: `" + adapter.TypeFile + "` Already Loaded!")
	}
	return fc
}

func SetMemory() {

}

func SetRedis() {

}

func SetMemcache() {

}
