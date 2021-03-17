package service

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/lbswl/academy-go-q12021/model"
)

// Reader reads data from a CSV file
func Reader() []model.Book {

	var books []model.Book

	f, err := os.Open("data/books.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Parse id to int
		id, errConv := strconv.Atoi(record[0])

		if errConv != nil {
			log.Fatal(errConv)
		}

		// Parse year to int
		year, errConv := strconv.ParseFloat(record[4], 64)

		if errConv != nil {
			log.Fatal(errConv)
		}

		books = append(books, model.Book{ID: id, Isbn: record[1], Authors: record[3], Year: int(year), ImageURL: record[5]})

	}

	return books
}

func Writer(crypto string) {
	file, err := os.OpenFile("data/data.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		log.Fatal("error while writing  data file")
	}

	cryptoString := []string{crypto}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{cryptoString}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}
