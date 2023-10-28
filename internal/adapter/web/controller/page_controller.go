package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
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
	if err := pc.renderer.Render("home.page.gohtml", w, nil); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Root(w http.ResponseWriter, r *http.Request) {
	if err := pc.renderer.Render("root.page.gohtml", w, nil); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) handleError(w http.ResponseWriter, httpErr payload.HTTPError) {
	if err := pc.renderer.Render("error.page.gohtml", w, httpErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
