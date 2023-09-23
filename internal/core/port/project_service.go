package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	ProjectService interface {
		Init(
			userId valueobject.ID,
			projectTitle string,
			groupTitle string,
		) (aggregate.Project,
			error)
	}
)
