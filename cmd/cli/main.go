package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/loader"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

func main() {
	var (
		configs = make(map[string]string)
		apiConf *config.Api
		err     error
	)

	//Get JSON configs from folder
	configs, err = loader.FilesToPaths(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	//Mapping configs to structs
	apiConf, err = loader.JSON[config.Api](configs["api_config.json"])
	if err != nil {
		logger.ERR.Fatalln(err)
	}
	logger.INFO.Println("Configs loaded")

	fmt.Println(apiConf)
}
