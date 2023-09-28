package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/generator"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"time"
)

type (
	projectService struct {
		projectRepo  port.ProjectRepo
		groupService port.GroupService
		userService  port.UserService
	}
)

func NewProjectService(
	projectRepo port.ProjectRepo,
	groupService port.GroupService,
	userService port.UserService,
) (port.ProjectService, error) {
	if projectRepo == nil {
		return nil, errorx.NilPointerErr{Item: "project repository"}
	}
	if groupService == nil {
		return nil, errorx.NilPointerErr{Item: "group service"}
	}
	if userService == nil {
		return nil, errorx.NilPointerErr{Item: "user service"}
	}
	return &projectService{
			projectRepo:  projectRepo,
			groupService: groupService,
			userService:  userService,
		},
		nil
}

func (ps projectService) Create(
	userId valueobject.UserID,
	projectTitle string,
	groupTitle string,
) (aggregate.Project, error) {
	var (
		user    aggregate.User
		project aggregate.Project
		err     error
	)

	user, err = ps.userService.Find(userId)
	if err != nil {
		return nil, err
	}

	project = aggregate.NewProjectBuilder(generator.UUID(), projectTitle).
		CreatedAt(time.Now()).
		UpdatedAt(time.Now()).
		Build()

	err = ps.projectRepo.Save(project)
	if err != nil {
		logger.ERR.Printf("Failed to save project (%s): %v\n", projectTitle, err)
		return nil, errorx.ProjectCreationErr{Title: projectTitle}
	}

	_, err = ps.groupService.CreateGroup(
		groupTitle, project.Id(),
		valueobject.Admin,
		user.Id(),
	)
	if err != nil {
		return nil, err
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
