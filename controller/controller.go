package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/lbswl/academy-go-q12021/model"

	"github.com/gorilla/mux"
)

type UseCase interface {
	FindUserById(Id int) ([]*model.UserJSON, error)
	ReadAllUsers() ([]*model.UserJSON, error)
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
		w.Write([]byte(`{"error": "Error reading the users file"}`))
		return
	}

	usersMarshalled, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the users file"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usersMarshalled)
}

// GetUser returns a user given its Id
func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, errConv := strconv.Atoi(params["id"])

	if errConv != nil {
		log.Fatal(errConv)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error reading the Id"}`))
	}

	user, err := c.useCase.FindUserById(id)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "User does not exist"}`))
		return
	}

	userMarshalled, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the users file"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userMarshalled)

}

//GetExternalData calls and external Id and saves the result to a CSV file
func (c *Controller) GetExternalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := c.useCase.GetExternalApiUsers()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error fecthing external data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "Fetched external data"}`))

}
