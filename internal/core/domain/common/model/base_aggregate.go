package model

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type BaseAggregate interface {
	Id() valueobject.ID
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
