package parser

import (
	"github.com/google/uuid"
)

func MustUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}

func UUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
