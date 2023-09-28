package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	GroupRepo interface {
		Save(group aggregate.Group) error
		Find(valueobject.GroupID) (aggregate.Group, error)
	}
)
