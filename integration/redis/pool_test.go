package redis

import (
	"github.com/SliverHorn/sliver/integration/redis/cache"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	expire := time.Second * 150 // 过期时间
	serializer := cache.SerializerJson // 序列化方式
	regularExpression := "^admins/\\d{1,5}$"
	expireNil := time.Second * 30 // 数据为空的过期时间
	c := NewCache(expire, serializer,regularExpression, expireNil)
	defer ReleaseCache(c)
}
