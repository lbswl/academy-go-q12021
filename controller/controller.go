package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lbswl/academy-go-q12021/config"
	"github.com/lbswl/academy-go-q12021/infrastructure/datastore"
	"github.com/lbswl/academy-go-q12021/model"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Get configutation variable
	config := config.GetConfig()
	books := datastore.Reader(config.CsvPath)
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a book given its ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	config := config.GetConfig()
	books := datastore.Reader(config.CsvPath)

	id, errConv := strconv.Atoi(params["id"])

	if errConv != nil {
		log.Fatal(errConv)
	}

	index := model.FindBookByID(books, id)

	if index > -1 {
		json.NewEncoder(w).Encode(books[index])
		return
	}
	json.NewEncoder(w).Encode(books)

}
