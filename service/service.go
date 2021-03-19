package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lbswl/academy-go-q12021/model"

	"github.com/gocarina/gocsv"
)

type Service interface {
	ReadFile() []*model.UserCSV
	WriteFile([]model.UserCSV)
	ClientExernalApi(numberCalls int)
}

type ServiceCSV struct {
	DataPath               string
	DataFile               string
	NumberCallsExternalApi int
	UrlExternalApi         string
}

func New(path string, file string, numCalls int, url string) ServiceCSV {
	return ServiceCSV{DataPath: path, NumberCallsExternalApi: numCalls,
		DataFile: file, UrlExternalApi: url}
}

func (s *ServiceCSV) ReadFile() ([]*model.UserCSV, error) {
	users := []*model.UserCSV{}

	fullPath := s.DataPath + s.DataFile
	usersFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return users, err
	}

	defer usersFile.Close()

	if err := gocsv.UnmarshalFile(usersFile, &users); err != nil { // Load clients from file
		return users, err
	}

	return users, nil

}

func (s *ServiceCSV) WriteFile(users []*model.UserCSV) error {
	fullPath := s.DataPath + s.DataFile
	usersFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer usersFile.Close()

	if _, err := usersFile.Seek(0, 0); err != nil { // Go to the end of the file
		return err
	}

	err = gocsv.MarshalFile(&users, usersFile) // Use this to save the CSV back to the file
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceCSV) ClientExernalApi() ([]*model.UserCSV, error) {

	users := []*model.UserCSV{}

	for i := 0; i < s.NumberCallsExternalApi; i++ {

		resp, err := http.Get(s.UrlExternalApi)

		if err != nil {
			return users, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return users, err
		}

		jsonData := []byte(string(body))

		var response model.Response

		err = json.Unmarshal(jsonData, &response)

		if err != nil {
			return users, err
		}

		users = append(users, &model.UserCSV{Id: i,
			Gender: response.Results[0].Gender, Title: response.Results[0].Name.Title,
			First: response.Results[0].Name.First, Last: response.Results[0].Name.Last,
			Email: response.Results[0].Email, CellPhone: response.Results[0].Cell, Nationality: response.Results[0].Nationality})

	}

	return users, nil
}
