package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
)

type projectController struct {
	*baseController
}

func NewProjectController() port.ProjectController {
	return &projectController{baseController: &baseController{}}
}
