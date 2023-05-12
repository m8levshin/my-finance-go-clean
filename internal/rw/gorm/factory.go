package gorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
