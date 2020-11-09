package interfaces

import "gorm.io/gorm"

type Gorm interface {
	Connect() *gorm.DB
	GormConfig() *gorm.Config
	AutoMigrateTables(db *gorm.DB)
}
