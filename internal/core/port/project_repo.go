package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	ProjectRepo interface {
		Save(project aggregate.Project) error
		Find(projectId valueobject.ID) (aggregate.Project, error)
	}
)
