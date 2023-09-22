package service

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"time"
)

type (
	projectService struct {
		projectRepo port.ProjectRepo
	}
)

func NewProjectService(projectRepo port.ProjectRepo) port.ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (ps *projectService) Create(title string) {
	var (
		project domain.Project
		err     error
	)
	project = domain.NewProjectBuilder(title).
		CreatedAt(time.Now()).
		UpdatedAt(time.Now()).
		Build()

	_, err = ps.projectRepo.Save(project)
	if err != nil {
		logger.ERR.Println(err)
		return
	}
}
