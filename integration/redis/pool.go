package redis

import (
	"github.com/SliverHorn/sliver/integration/redis/cache"
	"github.com/SliverHorn/sliver/integration/redis/operation"
	"github.com/SliverHorn/sliver/integration/redis/policy"
	"sync"
	"time"
)

var Pool *sync.Pool

func NewCache(expire time.Duration, serializer string, regularExpression string, expireNil time.Duration) *cache.Simple {
	Pool = &sync.Pool{New: func() interface {} {
		return cache.NewSimple(operation.NewString(), expire, serializer, policy.NewCrossPolicy(regularExpression, expireNil))
	}}
	return Pool.Get().(*cache.Simple)
}

func ReleaseCache(cache *cache.Simple) {
	Pool.Put(cache)
}