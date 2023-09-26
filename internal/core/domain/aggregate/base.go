package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	"time"
)

type (
	baseAggregate struct {
		id        model.ID
		createdAt time.Time
		updatedAt time.Time
	}
)

func newBaseAggregate(id model.ID) *baseAggregate {
	return &baseAggregate{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (b *baseAggregate) Id() model.ID {
	return b.id
}

func (b *baseAggregate) CreatedAt() time.Time {
	return b.createdAt
}

func (b *baseAggregate) UpdatedAt() time.Time {
	return b.updatedAt
}
