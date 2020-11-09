package boot

import (
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/model/admin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type Mysql struct{}

func (m *Mysql) Connect() *gorm.DB {
	var a = global.Config.Mysql
	dsn := a.Username + ":" + a.Password + "@tcp(" + a.Path + ")/" + a.Dbname + "?" + a.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), m.GormConfig()); err != nil {
		global.Zap.Error(global.I18n.TranslateFormat(`{#MysqlConnectFail} %v`, err))
		os.Exit(0)
		return nil
	} else {
		m.AutoMigrateTables(db)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(a.MaxIdleConns)
		sqlDB.SetMaxOpenConns(a.MaxOpenConns)
		return db
	}
}

func (m *Mysql) GormConfig() *gorm.Config {
	if global.Config.Mysql.LogMode {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}

func (m *Mysql) AutoMigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		new(admin.Admin),
	)
	if err != nil {
		global.Zap.Error(global.I18n.TranslateFormat(`{#RegisterTableFailed} %v`, err))
		os.Exit(0)
	}
	global.Zap.Info(global.I18n.TranslateFormat(`{#RegisterTableSuccess}`))
}

func InitGorm() {
	orm := &Mysql{}
	global.Db = orm.Connect()
}