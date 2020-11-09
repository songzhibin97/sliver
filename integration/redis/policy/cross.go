package policy

import (
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/redis/operation"
	"regexp"
	"time"
)

type CrossPolicy struct {
	// 正则表达式
	RegularExpression string

	// 过期时间
	Expire time.Duration

	// String
	operation *operation.String
}

func NewCrossPolicy(regularExpression string, expire time.Duration) *CrossPolicy {
	return &CrossPolicy{RegularExpression: regularExpression, Expire: expire}
}

func (c *CrossPolicy) IsNil(key string, value interface{}) {
	c.operation.Set(key, value, operation.WithExpiration(c.Expire)).UnWarp()
}

func (c *CrossPolicy) Before(key string) {
	if !regexp.MustCompile(c.RegularExpression).MatchString(key) {
		global.Zap.Error(global.I18n.Translate(`{#CrossPolicyFailed}`))
	}
}

func (c *CrossPolicy) SetOperation(operation *operation.String) {
	c.operation = operation
}

