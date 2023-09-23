package generator

import (
	"github.com/google/uuid"
)

func Uuid() uuid.UUID {
	return uuid.New()
}
