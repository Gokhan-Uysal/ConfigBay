package main

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/db"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/repo"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/service"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/loader"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"os"
)

var (
	configs = make(map[string]string)
	apiConf *config.Api
	dbConf  *config.Db
	err     error
)

func init() {
	//Get JSON configs from folder
	configs, err = loader.FilesToPaths(os.Getenv("CONF_PATH"))
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Mapping configs to structs
	apiConf, err = loader.JSON[config.Api](configs["api_config.json"])
	dbConf, err = loader.JSON[config.Db](configs["db_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	logger.INFO.Println("Configs loaded")
	fmt.Println(apiConf)
	fmt.Println(dbConf)
}

func main() {
	var (
		projectRepo port.ProjectRepo
		userRepo    port.UserRepo

		projectService port.ProjectService
		err            error
	)
	//Connect to db
	dsn := db.MakeDsn(dbConf)
	DB := db.Init("postgres", dsn)
	logger.INFO.Println("Db connected")
	_ = DB.Ping()

	//Init repos
	projectRepo, err = repo.NewProjectRepo(DB)
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	userRepo, err = repo.NewUserRepo(DB)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	//Init services
	projectService, err = service.NewProjectService(
		projectRepo,
		userRepo,
	)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	user := domain.NewUserBuilder(
		domain.NewUUID(),
		"john",
		domain.NewEmail("guysal20@ku.edu.tr"),
	).Build()

	_, err = userRepo.Create(user)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	project, err := projectService.Init(user.Id(), "Campus", "Admins")
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	logger.DEBUG.Println(project)
}
