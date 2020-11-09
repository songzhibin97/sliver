package result

import "github.com/SliverHorn/sliver/integration/redis/iterator"

type SliceResult interface {
	Iter(results []interface{}) *iterator.Iterator
	UnWarp() []interface{}
	UnWarpDefault(results []interface{}) []interface{}
}

type StringResult interface {
	UnWarp() string
	UnWarpElse(f func() string) string
	UnWarpDefault(d string) string
}

type InterfaceResult interface {
	UnWarp() interface{}
	UnWarpDefault(d interface{}) interface{}
}
