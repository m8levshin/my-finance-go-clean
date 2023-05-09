package domain

import "github.com/google/uuid"

type UpdatableProperty uint8

type Id uuid.UUID

func NewID() Id {
	return Id(uuid.New())
}
