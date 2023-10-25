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
	configs    = make(map[string]string)
	apiConf    *config.Api
	dbConf     *config.Db
	googleConf *config.Google
	err        error
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
	logger.INFO.Println("Configs loaded")
}

func main() {
	var (
		render         port.Renderer
		projectRepo    port.ProjectRepo
		groupRepo      port.GroupRepo
		userRepo       port.UserRepo
		userService    port.UserService
		groupService   port.GroupService
		_              port.ProjectService
		pageController port.PageController
		err            error
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
	_, err = service.NewProjectService(
		projectRepo,
		groupService,
		userService,
	)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Initialize controllers
	pageController, err = controller.NewPageController(render)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	handler := http.NewServeMux()
	staticPath := "/" + apiConf.Static
	handler.Handle(staticPath, http.StripPrefix(staticPath, fs))
	handler.Handle("/home", middleware.Get(http.HandlerFunc(pageController.Home)))

	url := fmt.Sprintf("%s:%s", apiConf.Host, strconv.Itoa(apiConf.Port))
	logger.ERR.Fatalln(http.ListenAndServe(url, handler))

}
