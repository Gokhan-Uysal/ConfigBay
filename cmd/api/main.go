package main

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/db"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/repo"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/service"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/generator"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/loader"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/mapper"
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
	configs, err = mapper.FilesToPaths(os.Getenv("CONF_PATH"))
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
		projectRepo    port.ProjectRepo
		groupRepo      port.GroupRepo
		userRepo       port.UserRepo
		userService    port.UserService
		groupService   port.GroupService
		projectService port.ProjectService
		err            error
	)

	//Connect to db
	dsn := db.MakeDsn(dbConf)
	DB := db.Init("postgres", dsn)
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

	user := aggregate.NewUserBuilder(
		generator.UUID(),
		"john",
		valueobject.NewEmail("guysal20@ku.edu.tr"),
	).Build()

	_, err = userRepo.Save(user)
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	project, err := projectService.Create(user.Id(), "Campus", "Admins")
	if err != nil {
		logger.ERR.Fatalln(err)
	}

	fmt.Println(project.Groups())
}
