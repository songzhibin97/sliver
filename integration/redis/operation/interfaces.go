package operation

import "github.com/SliverHorn/sliver/integration/redis/result"

type StringInterface interface {
	Set(key string, value interface{}, attrs ...*Attr) *result.Interface
	Get(key string) *result.String
	MGet(key ...string) *result.Slice
}