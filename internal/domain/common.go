package domain

import "github.com/google/uuid"

type Id uuid.UUID

func NewID() Id {
	return Id(uuid.New())
}
