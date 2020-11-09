package operation

import (
	"context"
	"github.com/SliverHorn/sliver/integration/redis/result"
	"time"
)

type String struct {
    ctx context.Context
}

func NewString() *String {
	return &String{ctx: context.Background()}
}

func (s *String) Set(key string, value interface{}, attrs ...*Attr) *result.Interface {
	 expiration := Attrs(attrs).Find(AttrExpiration).UnWarpDefault(time.Second*0).(time.Duration)
	 if nx := Attrs(attrs).Find(AttrNx).UnWarpDefault(nil); nx != nil {
	 	return result.NewInterface(Redis().SetNX(s.ctx, key, value, expiration).Result())
	 }
	 if xx := Attrs(attrs).Find(AttrXx).UnWarpDefault(nil); xx != nil {
	 	return result.NewInterface(Redis().SetXX(s.ctx, key, value, expiration).Result())
	 }
	return result.NewInterface(Redis().Set(s.ctx, key, value, expiration).Result())
}

func (s *String) Get(key string) *result.String {
	return result.NewString(Redis().Get(s.ctx, key).Result())
}

func (s *String) MGet(key ...string) *result.Slice {
	return result.NewSlice(Redis().MGet(s.ctx, key...).Result())
}

