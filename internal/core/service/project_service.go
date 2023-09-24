package service

import (
	aggregate2 "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	error2 "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/error"
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
	}
)

func NewProjectService(
	projectRepo port.ProjectRepo,
	groupRepo port.GroupRepo,
	userRepo port.UserRepo,
) (port.ProjectService, error) {
	if projectRepo == nil {
		return nil, error2.NilPointerErr{Item: "project repository"}
	}
	if groupRepo == nil {
		return nil, error2.NilPointerErr{Item: "group repository"}
	}
	if userRepo == nil {
		return nil, error2.NilPointerErr{Item: "user repository"}
	}
	return &projectService{projectRepo: projectRepo, groupRepo: groupRepo, userRepo: userRepo}, nil
}

func (ps projectService) Init(
	userId valueobject.ID,
	projectTitle string,
	groupTitle string,
) (aggregate2.Project,
	error) {
	var (
		user       aggregate2.User
		adminGroup aggregate2.Group
		project    aggregate2.Project
		err        error
	)

	user, err = ps.userRepo.Find(userId)
	if err != nil {
		logger.ERR.Printf("Failed to get user by ID (%s): %v\n", userId.String(), err)
		return nil, error2.UserNotFoundErr{Field: userId.String()}
	}

	adminGroup = aggregate2.NewGroupBuilder(generator.UUID(), groupTitle).
		Roles(
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
		return nil, error2.GroupCreationErr{Title: projectTitle}
	}

	project = aggregate2.NewProjectBuilder(generator.UUID(), projectTitle).
		CreatedAt(time.Now()).
		UpdatedAt(time.Now()).
		Groups(adminGroup.Id()).
		Build()

	err = ps.projectRepo.Save(project)
	if err != nil {
		logger.ERR.Printf("Failed to save project (%s): %v\n", projectTitle, err)
		return nil, error2.ProjectCreationErr{Title: projectTitle}
	}

	return project, nil
}

func (ps projectService) Find(projectId valueobject.ID) (aggregate2.Project, error) {
	var (
		project aggregate2.Project
		err     error
	)

	project, err = ps.projectRepo.Find(projectId)
	if err != nil {
		logger.ERR.Printf("Failed to find project (%s): %v\n", projectId.String(), err)
		return nil, error2.ProjectNotFoundErr{Id: projectId.String()}
	}

	return project, nil
}
