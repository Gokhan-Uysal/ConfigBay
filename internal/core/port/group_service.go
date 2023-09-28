package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	GroupService interface {
		CreateGroup(
			groupTitle string,
			projectId valueobject.ProjectID,
			role valueobject.Role,
			userIds ...valueobject.UserID,
		) (aggregate.Group, error)
	}
)
