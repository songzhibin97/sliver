package model

import (
	"github.com/SliverHorn/sliver/api/request"
	"gorm.io/gorm"
)

type Entity interface {
    Create(tx *gorm.DB, model interface{}) error
	Delete(tx *gorm.DB, id request.GetById) error
	Update(tx *gorm.DB, id request.GetById) error
	First(tx *gorm.DB, id request.GetById) error
	List(tx *gorm.DB, id request.PageInfo) error
}
