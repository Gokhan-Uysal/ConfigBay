package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/common"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"time"
)

type (
	projectService struct {
		projectRepo port.ProjectRepo
		userRepo    port.UserRepo
	}
)

func NewProjectService(
	projectRepo port.ProjectRepo,
	userRepo port.UserRepo,
) (port.ProjectService, error) {
	if projectRepo == nil {
		return nil, common.NilPointerErr{Item: "project repository"}
	}
	if userRepo == nil {
		return nil, common.NilPointerErr{Item: "user repository"}
	}
	return &projectService{projectRepo: projectRepo, userRepo: userRepo}, nil
}

func (ps *projectService) Init(
	userId domain.ID,
	projectTitle string,
	groupTitle string,
) (domain.Project,
	error) {
	var (
		user       domain.User
		adminGroup domain.Group
		project    domain.Project
		err        error
	)

	user, err = ps.userRepo.GetById(userId)
	if err != nil {
		logger.ERR.Printf("Failed to get user by ID (%s): %v\n", userId.String(), err)
		return nil, UserNotFoundErr{Field: userId.String()}
	}

	adminGroup = domain.NewGroupBuilder(domain.NewUUID(), groupTitle).
		Roles(
			domain.ManageGroups,
			domain.ManageUsers,
			domain.ReadSecrets,
			domain.WriteSecrets,
			domain.DeleteSecrets,
		).
		Users(user).
		Build()

	project = domain.NewProjectBuilder(domain.NewUUID(), projectTitle).
		CreatedAt(time.Now()).
		UpdatedAt(time.Now()).
		Groups(adminGroup).
		Build()

	err = ps.projectRepo.Save(project)
	if err != nil {
		logger.ERR.Printf("Failed to save project (%s): %v\n", projectTitle, err)
		return nil, ProjectCreationErr{Title: projectTitle}
	}

	return project, nil
}
