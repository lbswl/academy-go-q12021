package main

import (
	"log"
	"net/http"

	"github.com/lbswl/academy-go-q12021/controller"
	"github.com/lbswl/academy-go-q12021/router"
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

	serviceCSV := service.New(config.DataPath, config.DataFile,
		config.NumberCallsExternalApi, config.UrlExternalApi)
	useCase := usecase.New(serviceCSV)
	controller := controller.New(useCase)
	httpRouter := router.New(controller)

	log.Fatal(http.ListenAndServe("localhost:8000", httpRouter))

}
