package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
)

type pageController struct {
	*baseController
	renderer port.Renderer
}

func NewPageController(renderer port.Renderer) (port.PageController, error) {
	if renderer == nil {
		return nil, errorx.NilPointerErr{Item: "renderer"}
	}
	return pageController{baseController: &baseController{}, renderer: renderer}, nil
}

func (pc pageController) Favicon(w http.ResponseWriter, r *http.Request) {

}

func (pc pageController) Home(w http.ResponseWriter, r *http.Request) {
	if err := pc.renderer.Render("home.page.gohtml", w); err == nil {
		return
	}

	if err := pc.renderer.Render("404.page.gohtml", w); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
