package usecase

import (
	"encoding/json"
	"errors"

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
func (u *UseCase) ReadAllUsersConcurrently(params_type string, items int, items_per_workers int) ([]byte, error) {

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
