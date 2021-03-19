package main

import (
	"log"

	"github.com/lbswl/academy-go-q12021/util"
)

func main() {

	//Load configuration file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration file")
	}

	//Init Router
	//r := router.New()

	//log.Fatal(http.ListenAndServe(":8000", r))

}
