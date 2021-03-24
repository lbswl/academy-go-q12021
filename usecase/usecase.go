package usecase

import (
	"errors"

	"github.com/lbswl/academy-go-q12021/model"
	"github.com/lbswl/academy-go-q12021/service"
)

type UseCase struct {
	service service.ServiceCSV
}

func New(service service.ServiceCSV) *UseCase {
	return &UseCase{service}
}

func (u *UseCase) FindUserById(Id int) ([]*model.UserJSON, error) {
	usersJSON := []*model.UserJSON{}
	users, err := u.service.ReadFile()

	if err != nil {
		return usersJSON, err
	}

	for _, item := range users {
		if item.ID == Id {
			usersJSON = append(usersJSON, &model.UserJSON{ID: item.ID,
				Gender: item.Gender, Title: item.Title, First: item.First, Last: item.Last,
				Email: item.Email, CellPhone: item.CellPhone, Nationality: item.Nationality})
			return usersJSON, nil
		}
	}

	return usersJSON, errors.New("id not found")
}

func (u *UseCase) ReadAllUsers() ([]*model.UserJSON, error) {

	usersJSON := []*model.UserJSON{}
	usersCSV, err := u.service.ReadFile()

	if err != nil {
		return usersJSON, err
	}

	for _, item := range usersCSV {
		usersJSON = append(usersJSON, &model.UserJSON{ID: item.ID,
			Gender: item.Gender, Title: item.Title, First: item.First, Last: item.Last,
			Email: item.Email, CellPhone: item.CellPhone, Nationality: item.Nationality})
	}

	return usersJSON, nil

}

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
