package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
)

type projectController struct {
	handler http.Handler
}

func NewProjectController(handler http.Handler) port.ProjectController {
	return projectController{handler: handler}
}
