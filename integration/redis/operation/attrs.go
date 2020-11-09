package operation

import (
	"fmt"
	"github.com/SliverHorn/sliver/integration/redis/result"
	"time"
)

const (
	AttrNx = "nx"
	AttrXx = "xx"
	AttrExpiration = "expiration"
)

type empty struct {}

type Attr struct {
	Name  string
	Value interface{}
}

type Attrs []*Attr

func (a Attrs) Find(name string) *result.Interface {
	for _, attr := range a {
		if attr.Name == name {
			return result.NewInterface(attr.Value, nil)
		}
	}
	return result.NewInterface(nil, fmt.Errorf("attr not found name=%s failed", name))
}

func WithExpiration(t time.Duration) *Attr {
	return &Attr{Name: AttrExpiration, Value: t}
}

func WithNX() *Attr {
	return &Attr{Name: AttrNx, Value: empty{}}
}

func WithXX() *Attr {
	return &Attr{Name: AttrNx, Value: empty{}}
}
