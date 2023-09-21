package main

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/db"
	"os"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/loader"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

var (
	configs = make(map[string]string)
	apiConf *config.Api
	dbCOnf  *config.Db
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
	dbCOnf, err = loader.JSON[config.Db](configs["db_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	logger.INFO.Println("Configs loaded")
	fmt.Println(apiConf)
}

func main() {
	//Connect to db
	dsn := db.MakeDsn(dbCOnf)
	DB := db.Init("postgres", dsn)
	logger.INFO.Println("Db connected")
	_ = DB.Ping()
}
