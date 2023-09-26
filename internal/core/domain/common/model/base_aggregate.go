package model

import (
	"time"
)

type BaseAggregate interface {
	Id() ID
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
