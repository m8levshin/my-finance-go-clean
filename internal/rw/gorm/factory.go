package gorm

import (
	"github.com/mlevshin/my-finance-go-clean/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(configuration config.Configuration) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(configuration.Db.Dsn), &gorm.Config{})
}
