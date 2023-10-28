package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
	"time"
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
	homeItem := payload.NavbarItem{Href: "/home", Label: "Home"}
	projectItem := payload.NavbarItem{Href: "/projects", Label: "Projects"}

	campus := payload.ProjectItem{
		Icon:        make([]byte, 0),
		Title:       "Campus",
		Description: "App for all students",
		GroupCount:  5,
		UserCount:   10,
		CreatedAt:   time.Now(),
	}
	openWorld := payload.ProjectItem{
		Icon:        make([]byte, 0),
		Title:       "OpenWorld",
		Description: "App for all gamers",
		GroupCount:  2,
		UserCount:   20,
		CreatedAt:   time.Now(),
	}
	homePage := payload.HomePage{
		NavbarItems:  []payload.NavbarItem{homeItem, projectItem},
		ProjectItems: []payload.ProjectItem{campus, openWorld},
	}
	if err := pc.renderer.Render("home.page.gohtml", w, homePage); err == nil {
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
