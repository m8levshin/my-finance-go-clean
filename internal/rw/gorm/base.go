package gorm

import (
	"github.com/google/uuid"
)

type Base struct {
	ID *uuid.UUID `gorm:"type:uuid;primary_key;"`
}
