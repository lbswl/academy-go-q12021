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

// New returns an mux router
func New(controller Controller) *mux.Router {
	r := &mux.Router{}

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", controller.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{id}", controller.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/api/external", controller.GetExternalData).Methods(http.MethodGet)

	return r
}
