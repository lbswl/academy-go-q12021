package router

import (
	"github.com/lbswl/academy-go-q12021/controller"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := &mux.Router{}

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")

	return r
}
