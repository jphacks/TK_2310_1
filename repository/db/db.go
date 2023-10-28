package db

import "gorm.io/gorm"

type DB interface {
	Migrate() error
	Insert(model interface{}) error
	GetDB() *gorm.DB
}
