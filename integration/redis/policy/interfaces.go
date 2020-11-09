package policy

import "github.com/SliverHorn/sliver/integration/redis/operation"

// 缓存穿透
type CachePolicy interface {
	IsNil(key string, value interface{})
	Before(key string)
	SetOperation(operation *operation.String)
}
