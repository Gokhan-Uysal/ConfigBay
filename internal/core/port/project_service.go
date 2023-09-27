package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	ProjectService interface {
		Create(
			valueobject.UserID,
			string,
			string,
		) (aggregate.Project,
			error)
		Find(
			valueobject.ProjectID,
			valueobject.UserID,
		) (aggregate.Project, error)
	}
)
