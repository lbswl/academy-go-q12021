package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lbswl/academy-go-q12021/client"
	"github.com/lbswl/academy-go-q12021/model"
	"github.com/lbswl/academy-go-q12021/service"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := service.Reader()
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a book given its ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	books := service.Reader()

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

func GetExternalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	crypto, err := client.FetchCryptocurreyncy()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error fecthing external data"}`))
		return
	}

	//log.Printf("Crypto: %v", crypto)
	//Write crypto to the data file
	service.Writer(crypto)
	w.WriteHeader(http.StatusOK)

}
