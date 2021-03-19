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

func (u *UseCase) FindUserById(Id int) (*model.UserCSV, error) {
	users := u.service.ReadFile()

	for _, item := range users {
		if item.Id == Id {
			return item, nil
		}
	}

	return &model.UserCSV{}, errors.New("id not found")
}
