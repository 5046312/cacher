package cacher

import "github.com/5046312/cacher/adapter"

var (
	File     adapter.Adapter
	Memory   adapter.Adapter
	Redis    adapter.Adapter
	Memcache adapter.Adapter
)

// 每个adapter只能赋值一次
var adapters map[adapter.CacherType]adapter.Adapter

func FileConfig() *adapter.CacherFile {
	return adapter.DefaultCacherFileConfig()
}

func SetFile(fc *adapter.CacherFile) adapter.Adapter {
	// 判断是否已经加载过对应adapter
	if _, ok := adapters[adapter.TypeFile]; ok {
		panic("Cacher: `" + adapter.TypeFile + "` Already Loaded!")
	}
	adapters[adapter.TypeFile] = fc
	return fc
}

func SetMemory() {

}

func SetRedis() {

}

func SetMemcache() {

}
