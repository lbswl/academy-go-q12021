package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"runtime"
	"sync"

	"github.com/lbswl/academy-go-q12021/model"
	"github.com/lbswl/academy-go-q12021/service"
)

type UseCase struct {
	service service.ServiceCSV
}

// New returns a usecase struct
func New(service service.ServiceCSV) *UseCase {
	return &UseCase{service}
}

// FindUserbyId returns a user given its id
func (u *UseCase) FindUserById(Id int) ([]byte, error) {
	usersJSON := []*model.UserJSON{}
	users, err := u.service.ReadFile()

	if err != nil {
		return []byte(`{"error": "Error reading users file"}`), err
	}

	for _, item := range users {
		if item.ID == Id {
			usersJSON = append(usersJSON, &model.UserJSON{ID: item.ID,
				Gender: item.Gender, Title: item.Title, First: item.First, Last: item.Last,
				Email: item.Email, CellPhone: item.CellPhone, Nationality: item.Nationality})
			userMarshalled, err := json.Marshal(usersJSON)

			if err != nil {
				return []byte(`{"error": "Error marshalling users file"}`), err
			}

			return userMarshalled, nil
		}
	}

	return []byte(`{"error": "User does not exist"}`), errors.New("id not found")
}

// ReadAllUsers returns all users in the csv file
func (u *UseCase) ReadAllUsers() ([]byte, error) {

	usersJSON := []*model.UserJSON{}
	usersCSV, err := u.service.ReadFile()

	if err != nil {
		return []byte(`{"error": "Error reading the users file"}`), err
	}

	for _, item := range usersCSV {
		usersJSON = append(usersJSON, &model.UserJSON{ID: item.ID,
			Gender: item.Gender, Title: item.Title, First: item.First, Last: item.Last,
			Email: item.Email, CellPhone: item.CellPhone, Nationality: item.Nationality})
	}

	usersMarshalled, err := json.Marshal(usersJSON)

	if err != nil {
		return []byte(`{"error": "Error marshalling the users file"}`), err
	}

	return usersMarshalled, nil

}

// GetExternalApiUsers writes the response from external API to users csv file
func (u *UseCase) GetExternalApiUsers() error {

	userCSV, err := u.service.ClientExernalApi()
	if err != nil {
		return err
	}

	err = u.service.WriteFile(userCSV)
	if err != nil {
		return err
	}

	return nil

}

// ReadAllUsersConcurrently returns all users in the csv file
func (u *UseCase) ReadAllUsersConcurrently(paramsType string, items int, itemsPerWorkers int) ([]byte, error) {

	usersCSV, err := u.service.ReadFile()

	if err != nil {
		return []byte(`{"error": "Error reading the users file"}`), err
	}

	usersJSON := []*model.UserJSON{}

	values := make(chan int)
	shutdown := make(chan struct{})
	poolSize := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(id int, itemsPerWorkers int) {
			n := id
			currentNumItems := 0
			for {

				select {
				case values <- n:
					log.Printf("Worker %d sent %d\n", id, n)
				case <-shutdown:
					log.Printf("Worker %d shutting down\n", id)
					wg.Done()
					return
				}

				currentNumItems++
				if currentNumItems < itemsPerWorkers {
					n = n + poolSize
				}
			}
		}(i, itemsPerWorkers)
	}

	for i := range values {

		if paramsType == "odd" && i%2 == 0 {
			continue
		}

		if paramsType == "even" && i%2 != 0 {
			continue
		}

		if len(usersJSON) == items || len(usersJSON) == len(usersCSV) || i > len(usersCSV)-1 {
			break
		}

		usersJSON = append(usersJSON, &model.UserJSON{ID: usersCSV[i].ID,
			Gender: usersCSV[i].Gender, Title: usersCSV[i].Title, First: usersCSV[i].First, Last: usersCSV[i].Last,
			Email: usersCSV[i].Email, CellPhone: usersCSV[i].CellPhone, Nationality: usersCSV[i].Nationality})
	}

	close(shutdown)
	wg.Wait()

	usersMarshalled, err := json.Marshal(usersJSON)

	if err != nil {
		return []byte(`{"error": "Error marshalling the users file"}`), err
	}

	return usersMarshalled, nil

}
