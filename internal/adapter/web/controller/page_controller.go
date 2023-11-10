package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/pagedata"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/parser"
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
	rootPage := pagedata.Root{
		Config: pc.rootPageConf,
	}
	if err := pc.renderer.Render("root.page.gohtml", w, rootPage); err == nil {
		return
	}

	pc.handleErrorPage(w, payload.InternalServerErr)
}

func (pc pageController) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("code")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	project, err := pc.projectService.Find(
		parser.MustUUID("aa8e365a-3297-4c9c-a217-890a2c1eef06"),
		parser.MustUUID("647e3367-e764-4088-9031-64a74a4fec3c"),
	)
	if err != nil {
		return
	}

	homePage := pagedata.HomePage{
		Config: pc.homePageConf,
		ProjectItems: []pagedata.ProjectItem{
			pagedata.ToProjectItem(project),
		},
	}

	if err := pc.renderer.Render("home.page.gohtml", w, homePage); err == nil {
		return
	}
	pc.handleErrorPage(w, payload.InternalServerErr)
}

func (pc pageController) Signup(w http.ResponseWriter, r *http.Request) {
	onboardPage := pagedata.Onboard{
		Config: pc.onboardPageConf,
		Access: pagedata.Signup,
	}
	if err := pc.renderer.Render("onboard.page.gohtml", w, onboardPage); err == nil {
		return
	}

	pc.handleErrorPage(w, payload.InternalServerErr)
}

func (pc pageController) Login(w http.ResponseWriter, r *http.Request) {
	onboardPage := pagedata.Onboard{
		Config: pc.onboardPageConf,
		Access: pagedata.Login,
	}
	if err := pc.renderer.Render("onboard.page.gohtml", w, onboardPage); err == nil {
		return
	}

	pc.handleErrorPage(w, payload.InternalServerErr)
}

func (pc pageController) handleErrorPage(w http.ResponseWriter, httpErr payload.HTTPError) {
	if err := pc.renderer.Render("error.page.gohtml", w, httpErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
