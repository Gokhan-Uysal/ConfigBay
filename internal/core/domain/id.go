package domain

import "github.com/google/uuid"

type ID interface {
	String() string
}

func NewUUID() ID {
	return uuid.New()
}

func ParseToUUID(id string) (ID, error) {
	return uuid.Parse(id)
}
