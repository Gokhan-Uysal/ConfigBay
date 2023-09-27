package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

type (
	userService struct {
		userRepo port.UserRepo
	}
)

func NewUserService(
	userRepo port.UserRepo,
) (port.UserService, error) {
	if userRepo == nil {
		return nil, errorx.NilPointerErr{Item: "user repository"}
	}
	return &userService{
			userRepo: userRepo,
		},
		nil
}

func (us userService) Find(userId valueobject.UserID) (aggregate.User, error) {
	var (
		user aggregate.User
		err  error
	)

	user, err = us.userRepo.Find(userId)
	if err != nil {
		logger.ERR.Printf("Failed to get user by ID (%s): %v\n", userId.String(), err)
		return nil, errorx.UserNotFoundErr{Field: userId.String()}
	}

	return user, nil
}
