package service

import (
	"fmt"
	"github.com/SliverHorn/sliver/api/request"
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/interfaces"
	"github.com/SliverHorn/sliver/model/admin"
	"github.com/SliverHorn/sliver/integration/redis"
	"github.com/SliverHorn/sliver/integration/redis/cache"
	"gorm.io/gorm"
	"time"
)

type Admin struct{}

func (a *Admin) Create(tx *gorm.DB, model interface{}) (err error) {
	var adminToDb admin.Admin
	if data, ok := model.(admin.Admin); ok {
		adminToDb = data
		s := interfaces.NewSnowflake()
		if adminToDb.Id, err = s.GenId(); err != nil {
			return err
		}
	}
	defer func() {
		if recover() != nil {
			tx.Rollback()
		}
	}()
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&adminToDb).Error; err != nil {
			return err
		}
		return nil
	})
}

func (a *Admin) Delete(tx *gorm.DB, id request.GetById) (err error) {
	return
}

func (a *Admin) Update(tx *gorm.DB, id request.GetById) (err error) {
	return
}

func (a *Admin) First(tx *gorm.DB, id request.GetById) (err error) {

	return
}

func (a *Admin) List(tx *gorm.DB, id request.PageInfo) (err error) {
	return
}

func FindAdmin(id string) *admin.Admin {
	expire := time.Second * 150 // 过期时间
	serializer := cache.SerializerJson // 序列化方式
	regularExpression := "^admins/\\d{1,5}$"
	expireNil := time.Second * 30 // 数据为空的过期时间
	c := redis.NewCache(expire, serializer,regularExpression, expireNil)
	defer redis.ReleaseCache(c)
	c.DbGetter = NewAdminIdGetter(id)
	a := admin.NewAdmin()
	c.GetModel("admins/1", a)
	return a
}

func NewAdminIdGetter(id string) cache.DbGetter {
	return func() interface{} {
		fmt.Println("Get Data From Db")
		var a admin.Admin
		if err := global.Db.Debug().Where("id=?", id).First(&a).Error; err != nil {
			global.Zap.Error(global.I18n.TranslateFormat(`{#IdGetterError} %v`, err))
			return a
		}
		return a
	}
}
