package parser

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

func ToUuid(id string) (valueobject.ID, error) {
	return uuid.Parse(id)
}
