package global

import (
	"github.com/SliverHorn/sliver/api/config"
	"github.com/gogf/gf/i18n/gi18n"
	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Zap    *zap.Logger
	I18n   *gi18n.Manager
	Redis  *redis.Client
	Viper  *viper.Viper
	Config config.Server
)
