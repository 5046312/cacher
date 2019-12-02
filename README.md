# cacher
my cache tool for golang


## File

在需要操作interface{}的场景，需要程序先执行过该interface{}的保存方法后，才能进行读取。
这是因为做了gob的转码处理，在保存方法中才进行gob的注册。
如果没有执行过保存而是直接读取缓存，会获取nil
