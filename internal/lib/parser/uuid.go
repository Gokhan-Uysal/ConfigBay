package parser

import (
	"github.com/google/uuid"
)

func ToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
