package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UseCase interface {
	FindUserById(Id int) ([]byte, error)
	ReadAllUsers() ([]byte, error)
	ReadAllUsersConcurrently(params_type string, items int, items_per_workers int) ([]byte, error)
	GetExternalApiUsers() error
}

type Controller struct {
	useCase UseCase
}

func New(u UseCase) *Controller {
	return &Controller{u}
}

// GetUsers returns all users
func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := c.useCase.ReadAllUsers()

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(users)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

// GetUser returns a user given its Id
func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, errConv := strconv.Atoi(params["id"])

	if errConv != nil {
		log.Fatal(errConv)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the Id parameter"}`))
	}

	user, err := c.useCase.FindUserById(id)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(user)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(user)

}

//GetExternalData calls and external Id and saves the result to a CSV file
func (c *Controller) GetExternalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := c.useCase.GetExternalApiUsers()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error fecthing external data"}`))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "Fetched external data"}`))

}

// GetUsersConturrent returns all users (reads the using concurrency)
func (c *Controller) GetUsersConcurrent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	params_type := params["type"]

	items, errConv := strconv.Atoi(params["items"])

	if errConv != nil {
		log.Fatal(errConv)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the items parameter"}`))
	}

	items_per_workers, errConv := strconv.Atoi(params["items_per_workers"])

	if errConv != nil {
		log.Fatal(errConv)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error parsing the items_per_workers parameter"}`))
	}

	users, err := c.useCase.ReadAllUsersConcurrently(params_type, items, items_per_workers)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(users)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(users)
}
