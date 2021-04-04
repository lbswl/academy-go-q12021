package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lbswl/academy-go-q12021/controller"
	"github.com/lbswl/academy-go-q12021/router"
	"github.com/lbswl/academy-go-q12021/service"
	"github.com/lbswl/academy-go-q12021/usecase"
	"github.com/lbswl/academy-go-q12021/util"
)

// Abnormal exit constants
const (
	ExitAbnormalErrorLoadingConfiguration = iota
	ExitAbnormalErrorLoadingCSVFile
)

func main() {

	//Load configuration file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration file")
	}

	fullPath := config.DataPath + config.DataFile

	rf, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}

	wf, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		os.Exit(ExitAbnormalErrorLoadingCSVFile)
	}

	defer rf.Close()
	defer wf.Close()

	serviceCSV := service.New(rf, wf, config.NumberCallsExternalApi, config.UrlExternalApi)
	useCase := usecase.New(serviceCSV)
	controller := controller.New(useCase)
	httpRouter := router.New(controller)

	log.Fatal(http.ListenAndServe("localhost:8000", httpRouter))

}
