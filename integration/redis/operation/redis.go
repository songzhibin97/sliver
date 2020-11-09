package operation

import (
	"context"
	"github.com/SliverHorn/sliver/global"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

func Redis() (client *redis.Client){
	var once sync.Once
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			DB:       global.Config.Redis.DB, // use default DB
			Addr:     global.Config.Redis.Addr,
			Password: global.Config.Redis.Password, // no password set

			PoolSize:     15, // 连接池
			MinIdleConns: 10, // 最小连接数
			// 超时
			DialTimeout:  5 * time.Second, // 连接建立超时时间
			ReadTimeout:  3 * time.Second, // 读超时,默认3秒,-1代表取消读超时
			WriteTimeout: 3 * time.Second, // 写超时,默认等于读超时
			PoolTimeout:  4 * time.Second, // 当所有连接都处于繁忙状态时,客户端等待可用连接的最大等待时长
			//
			MaxRetries: 0, // 放弃之前的最大重试次数,默认是不重试失败的命令。
			MinRetryBackoff: 8 * time.Millisecond, // 每次重试之间的最小退避,默认值为8毫秒,-1禁用退避。
			MaxRetryBackoff: 512 * time.Millisecond, // 每次重试之间的最大退避,默认值为512毫秒,-1禁用退避。
		})

		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			global.Zap.Info(global.I18n.TranslateFormat("{#PingFailed} %v", err))
		} else {
			global.Zap.Info(global.I18n.TranslateFormat("{#PingResponse} %v", pong))
		}
	})
	return client
}
