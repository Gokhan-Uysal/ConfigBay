package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	baseAggregate struct {
		id        valueobject.ID
		createdAt time.Time
		updatedAt time.Time
	}
)

func newBaseAggregate(id valueobject.ID) *baseAggregate {
	return &baseAggregate{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (b *baseAggregate) Id() valueobject.ID {
	return b.id
}

func (b *baseAggregate) CreatedAt() time.Time {
	return b.createdAt
}

func (b *baseAggregate) UpdatedAt() time.Time {
	return b.updatedAt
}
