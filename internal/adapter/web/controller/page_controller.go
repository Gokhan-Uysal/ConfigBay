package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
	"time"
)

type pageController struct {
	*baseController
	renderer       port.Renderer
	ssoProviders   []config.SSOProvider
	projectService port.ProjectService
}

func NewPageController(
	renderer port.Renderer,
	ssoProviders []config.SSOProvider,
	projectService port.ProjectService,
) (port.PageController, error) {
	if renderer == nil {
		return nil, errorx.NilPointerErr{Item: "renderer"}
	}
	if ssoProviders == nil {
		return nil, errorx.NilPointerErr{Item: "sso providers"}
	}
	if projectService == nil {
		return nil, errorx.NilPointerErr{Item: "project service"}
	}
	return pageController{
		baseController: &baseController{},
		renderer:       renderer,
		ssoProviders:   ssoProviders,
		projectService: projectService,
	}, nil
}

func (pc pageController) Favicon(w http.ResponseWriter, r *http.Request) {

}

func (pc pageController) Root(w http.ResponseWriter, r *http.Request) {
	rootPage := payload.RootPage{NavbarItems: payload.RootPageNavbar}
	if err := pc.renderer.Render("root.page.gohtml", w, rootPage); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Home(w http.ResponseWriter, r *http.Request) {
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
		NavbarItems:  payload.HomePageNavbar,
		ProjectItems: []payload.ProjectItem{campus, openWorld},
	}
	if err := pc.renderer.Render("home.page.gohtml", w, homePage); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Signup(w http.ResponseWriter, r *http.Request) {
	onboardPage := payload.OnboardPage{
		NavbarItems:  payload.RootPageNavbar,
		Access:       "signup",
		OnboardItems: []payload.OnboardItem{payload.GoogleItem, payload.GithubItem},
	}

	if err := pc.renderer.Render("onboard.page.gohtml", w, onboardPage); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Login(w http.ResponseWriter, r *http.Request) {
	onboardPage := payload.OnboardPage{
		NavbarItems:  payload.RootPageNavbar,
		Access:       "login",
		OnboardItems: []payload.OnboardItem{payload.GoogleItem, payload.GithubItem},
	}

	if err := pc.renderer.Render("onboard.page.gohtml", w, onboardPage); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) handleError(w http.ResponseWriter, httpErr payload.HTTPError) {
	if err := pc.renderer.Render("error.page.gohtml", w, httpErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
