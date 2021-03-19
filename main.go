package main

import (
	"log"

	"github.com/lbswl/academy-go-q12021/service"
	"github.com/lbswl/academy-go-q12021/usecase"
	"github.com/lbswl/academy-go-q12021/util"
)

func main() {

	//Load configuration file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration file")
	}

	serviceCSV := service.New(config.DataPath, config.DataFile, config.UrlExternalApi)
	useCase := usecase.New(serviceCSV)

	//Init Router
	//r := router.New()

	//log.Fatal(http.ListenAndServe(":8000", r))

}
