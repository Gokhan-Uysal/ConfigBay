package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/generator"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"time"
)

type (
	projectService struct {
		projectRepo port.ProjectRepo
		groupRepo   port.GroupRepo
		userRepo    port.UserRepo
		authService port.AuthService
	}
)

func NewProjectService(
	projectRepo port.ProjectRepo,
	groupRepo port.GroupRepo,
	userRepo port.UserRepo,
	authService port.AuthService,
) (port.ProjectService, error) {
	if projectRepo == nil {
		return nil, errorx.NilPointerErr{Item: "project repository"}
	}
	if groupRepo == nil {
		return nil, errorx.NilPointerErr{Item: "group repository"}
	}
	if userRepo == nil {
		return nil, errorx.NilPointerErr{Item: "user repository"}
	}
	if authService == nil {
		return nil, errorx.NilPointerErr{Item: "authentication service"}
	}
	return &projectService{
			projectRepo: projectRepo,
			groupRepo:   groupRepo,
			userRepo:    userRepo,
			authService: authService,
		},
		nil
}

func (ps projectService) Init(
	userId valueobject.UserID,
	projectTitle string,
	groupTitle string,
) (aggregate.Project, error) {
	var (
		user       aggregate.User
		adminGroup aggregate.Group
		project    aggregate.Project
		err        error
	)

	user, err = ps.userRepo.Find(userId)
	if err != nil {
		logger.ERR.Printf("Failed to get user by ID (%s): %v\n", userId.String(), err)
		return nil, errorx.UserNotFoundErr{Field: userId.String()}
	}

	adminGroup = aggregate.NewGroupBuilder(generator.UUID(), groupTitle).
		Roles(
			entity.ReadProject,
			entity.ManageGroups,
			entity.ManageUsers,
			entity.ReadSecrets,
			entity.WriteSecrets,
			entity.DeleteSecrets,
		).
		Users(user.Id()).
		Build()

	err = ps.groupRepo.Save(adminGroup)
	if err != nil {
		logger.ERR.Printf("Failed to save group (%s): %v\n", groupTitle, err)
		return nil, errorx.GroupCreationErr{Title: projectTitle}
	}

	project = aggregate.NewProjectBuilder(generator.UUID(), projectTitle).
		CreatedAt(time.Now()).
		UpdatedAt(time.Now()).
		Groups(adminGroup.Id()).
		Build()

	err = ps.projectRepo.Save(project)
	if err != nil {
		logger.ERR.Printf("Failed to save project (%s): %v\n", projectTitle, err)
		return nil, errorx.ProjectCreationErr{Title: projectTitle}
	}

	return project, nil
}

func (ps projectService) Find(
	projectId valueobject.ProjectID,
	userId valueobject.UserID,
) (aggregate.Project, error) {
	var (
		project aggregate.Project
		err     error
	)

	project, err = ps.projectRepo.Find(projectId)
	if err != nil {
		logger.ERR.Printf("Failed to find project (%s): %v\n", projectId.String(), err)
		return nil, errorx.ProjectNotFoundErr{Id: projectId.String()}
	}

	return project, nil
}
