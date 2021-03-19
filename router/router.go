package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetExternalData(w http.ResponseWriter, r *http.Request)
}

func New(controller Controller) *mux.Router {
	r := &mux.Router{}

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", controller.GetUserById).Methods("GET")
	r.HandleFunc("/api/external", controller.GetExternalData).Methods("GET")

	return r
}
