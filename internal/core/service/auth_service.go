package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"slices"
)

type (
	authService struct {
		groupRepo port.GroupRepo
		userRepo  port.UserRepo
	}
)

func NewAuthService() (port.AuthService, error) {
	return &authService{},
		nil
}

func (as authService) HasRoles(groupId valueobject.GroupID, roles ...entity.Role) bool {
	var (
		group aggregate.Group
		err   error
	)

	group, err = as.groupRepo.Find(groupId)
	if err != nil {
		return false
	}

	for _, role := range roles {
		if !slices.Contains(group.Roles(), role) {
			return false
		}
	}

	return true
}
