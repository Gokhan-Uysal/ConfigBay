package controller

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
)

type pageController struct {
	*baseController
	renderer        port.Renderer
	ssoProviders    []config.SSOProvider
	rootPageConf    *config.RootPage
	homePageConf    *config.HomePage
	onboardPageConf *config.OnboardPage
	projectService  port.ProjectService
}

func NewPageController(
	renderer port.Renderer,
	ssoProviders []config.SSOProvider,
	rootPageConf *config.RootPage,
	homePageConf *config.HomePage,
	onboardPageConf *config.OnboardPage,
	projectService port.ProjectService,
) (port.PageController, error) {
	if renderer == nil {
		return nil, errorx.NilPointerErr{Item: "renderer"}
	}
	if ssoProviders == nil {
		return nil, errorx.NilPointerErr{Item: "sso providers"}
	}
	if rootPageConf == nil || homePageConf == nil || onboardPageConf == nil {
		return nil, errorx.NilPointerErr{Item: "page configurations"}
	}
	if projectService == nil {
		return nil, errorx.NilPointerErr{Item: "project service"}
	}

	fmt.Println(onboardPageConf)
	return pageController{
		baseController:  &baseController{},
		renderer:        renderer,
		ssoProviders:    ssoProviders,
		rootPageConf:    rootPageConf,
		homePageConf:    homePageConf,
		onboardPageConf: onboardPageConf,
		projectService:  projectService,
	}, nil
}

func (pc pageController) Favicon(w http.ResponseWriter, r *http.Request) {

}

func (pc pageController) Root(w http.ResponseWriter, r *http.Request) {
	if err := pc.renderer.Render("root.page.gohtml", w, pc.rootPageConf); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Home(w http.ResponseWriter, r *http.Request) {
	if err := pc.renderer.Render("home.page.gohtml", w, pc.homePageConf); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Signup(w http.ResponseWriter, r *http.Request) {
	pc.onboardPageConf.Name = "Signup"
	if err := pc.renderer.Render("onboard.page.gohtml", w, pc.onboardPageConf); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) Login(w http.ResponseWriter, r *http.Request) {
	pc.onboardPageConf.Name = "Login"
	if err := pc.renderer.Render("onboard.page.gohtml", w, pc.onboardPageConf); err == nil {
		return
	}

	pc.handleError(w, payload.InternalServerErr)
}

func (pc pageController) handleError(w http.ResponseWriter, httpErr payload.HTTPError) {
	if err := pc.renderer.Render("error.page.gohtml", w, httpErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
