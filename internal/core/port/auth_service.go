package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	AuthService interface {
		HasRoles(valueobject.GroupID, ...entity.Role) bool
	}
)
