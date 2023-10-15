package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/generator"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

type (
	groupService struct {
		groupRepo port.GroupRepo
	}
)

func NewGroupService(
	groupRepo port.GroupRepo,
) (port.GroupService, error) {
	if groupRepo == nil {
		return nil, errorx.NilPointerErr{Item: "group repository"}
	}
	return &groupService{
			groupRepo: groupRepo,
		},
		nil
}

func (gs groupService) CreateGroup(
	groupTitle string,
	projectId valueobject.ProjectID,
	role valueobject.Role,
	userIds ...valueobject.UserID,
) (aggregate.Group, error) {
	var (
		adminGroup aggregate.Group
		err        error
	)
	adminGroup = aggregate.NewGroupBuilder(generator.UUID(), groupTitle, projectId).
		Role(role).
		Users(userIds...).
		Build()

	err = gs.groupRepo.Save(adminGroup)
	if err != nil {
		logger.ERR.Printf("Failed to save group (%s): %v\n", groupTitle, err)
		return nil, errorx.GroupCreationErr{Title: groupTitle}
	}

	return adminGroup, nil
}
