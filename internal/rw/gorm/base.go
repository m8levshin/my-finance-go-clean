package gorm

import (
	"github.com/google/uuid"
)

type Base struct {
	Id uuid.UUID `gorm:"type:uuid;primary_key;"`
}
