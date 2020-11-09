package cache

// 缓存接口
type SimpleCache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	GetModel(key string, model interface{})
}
