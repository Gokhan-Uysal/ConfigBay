package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
)

type pageController struct {
	*baseController
	renderer port.Renderer
}

func NewPageController(renderer port.Renderer) port.PageController {
	return pageController{baseController: &baseController{}}
}

func (pc pageController) Home(w http.ResponseWriter, r *http.Request) {
	if err := pc.renderer.Render("home.page.gohtml", w); err == nil {
		return
	}

	if err := pc.renderer.Render("404.page.gohtml", w); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
