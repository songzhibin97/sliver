package cache

import (
	"fmt"
	"github.com/SliverHorn/sliver/integration/redis/operation"
	"testing"
	"time"
)

func TestNewSimple(t *testing.T) {
	cache := NewSimple(operation.NewString(), time.Second*17)
	cache.DbGetter = func() string {
		return "data comes from db"
	}
	fmt.Println(cache.Get("123"))
}
