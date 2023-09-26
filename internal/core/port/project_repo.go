package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	ProjectRepo interface {
		Save(aggregate.Project) error
		Find(valueobject.ProjectID) (aggregate.Project, error)
	}
)
