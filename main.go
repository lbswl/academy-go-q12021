package main

import (
	"log"
	"net/http"

	"github.com/lbswl/academy-go-q12021/router"
)

func main() {

	//Init Router
	r := router.New()

	log.Fatal(http.ListenAndServe(":8000", r))

}
