package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	GroupService interface {
		Create(
			string,
			valueobject.ProjectID,
			valueobject.UserID,
			...valueobject.Role,
		) (aggregate.Group, error)
	}
)
