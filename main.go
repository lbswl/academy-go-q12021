package main

import (
	"log"
	"net/http"

	"github.com/lbswl/academy-go-q12021/controller"

	"github.com/gorilla/mux"
)

func main() {

	//Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
