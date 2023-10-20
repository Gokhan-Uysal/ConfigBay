package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
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

func (pc pageController) Home(rw http.ResponseWriter, r *http.Request) {
	var (
		httpErr *payload.HTTPError
	)

	if r.Method != http.MethodGet {
		httpErr = &payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: "Page Not Found!",
		}
		rw.WriteHeader(http.StatusNotFound)

		if err := pc.renderer.Render("error.page.gohtml", rw, httpErr); err != nil {
			logger.ERR.Println(err)
		}
		return
	}

	if err := pc.renderer.Render("home.page.gohtml", rw, nil); err == nil {
		return
	}

	httpErr = &payload.HTTPError{
		StatusCode: http.StatusInternalServerError, StatusMessage: "Boom!",
	}
	if err := pc.renderer.Render("error.page.gohtml", rw, httpErr); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
