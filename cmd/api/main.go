package main

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/db"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/repo"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/controller"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/middleware"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/renderer"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/service"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/loader"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/mapper"
	"net/http"
	"os"
	"strconv"
)

var (
	configs         map[string]string
	apiConf         *config.Api
	dbConf          *config.Db
	googleConf      *config.Google
	rootPageConf    *config.RootPage
	homePageConf    *config.HomePage
	onboardPageConf *config.OnboardPage
	err             error
)

func init() {
	//Get JSON configs from folder
	configs, err = mapper.FilesToPaths(os.Getenv("CONF_PATH"))
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Mapping configs to structs
	apiConf, err = loader.JSON[config.Api](configs["api_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	dbConf, err = loader.JSON[config.Db](configs["db_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	googleConf, err = loader.JSON[config.Google](configs["google_sso_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	rootPageConf, err = loader.JSON[config.RootPage](configs["root_page_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	homePageConf, err = loader.JSON[config.HomePage](configs["home_page_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	onboardPageConf, err = loader.JSON[config.OnboardPage](configs["onboard_page_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	logger.INFO.Println("Configs loaded")
}

func main() {
	var (
		render            port.Renderer
		projectRepo       port.ProjectRepo
		groupRepo         port.GroupRepo
		userRepo          port.UserRepo
		userService       port.UserService
		groupService      port.GroupService
		projectService    port.ProjectService
		pageController    port.PageController
		onboardController port.OnboardController
		err               error
	)

	//Generate html template cache
	render = renderer.New()
	err = render.Load(apiConf.Template)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	logger.INFO.Println("Template cache created")

	//Link css and javascript files
	fs := http.FileServer(http.Dir(apiConf.Static))
	logger.INFO.Println("File server created")

	//Connect to db
	dsn := db.MakeDsn(dbConf)
	DB := db.Init(dbConf.Driver, dsn)
	logger.INFO.Println("Db connected")
	_ = DB.Ping()

	//Initialize repositories
	projectRepo, err = repo.NewProjectRepo(DB)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	groupRepo, err = repo.NewGroupRepo(DB)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	userRepo, err = repo.NewUserRepo(DB)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Initialize services
	userService, err = service.NewUserService(userRepo)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	groupService, err = service.NewGroupService(groupRepo)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	projectService, err = service.NewProjectService(
		projectRepo,
		groupService,
		userService,
	)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Initialize controllers
	pageController, err = controller.NewPageController(
		render,
		apiConf.SSOProviders,
		rootPageConf,
		homePageConf,
		onboardPageConf,
		projectService,
	)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	onboardController, err = controller.NewOnboardController(googleConf)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	handler := http.NewServeMux()
	staticPath := config.Root.String() + apiConf.Static

	handler.Handle(staticPath, http.StripPrefix(staticPath, fs))
	handler.Handle(config.Root.String(), middleware.Get(http.HandlerFunc(pageController.Root)))
	handler.Handle(config.Home.String(), middleware.Get(http.HandlerFunc(pageController.Home)))
	handler.Handle(config.Signup.String(), middleware.Get(http.HandlerFunc(pageController.Signup)))
	handler.Handle(config.Login.String(), middleware.Get(http.HandlerFunc(pageController.Login)))

	handler.Handle(
		config.SignupWith.String(),
		middleware.Get(
			http.HandlerFunc(
				onboardController.SignupWith,
			),
		),
	)
	handler.Handle(
		config.LoginWith.String(), middleware.Get(http.HandlerFunc(onboardController.LoginWith)),
	)
	handler.Handle(
		config.RedirectGoogle.String(),
		middleware.Get(
			http.HandlerFunc(
				onboardController.RedirectGoogle,
			),
		),
	)

	url := fmt.Sprintf("%s:%s", apiConf.Host, strconv.Itoa(apiConf.Port))
	logger.INFO.Printf("Server is listening on %s\n", url)
	logger.ERR.Fatalln(http.ListenAndServe(url, handler))

}
