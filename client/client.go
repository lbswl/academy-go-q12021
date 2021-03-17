package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lbswl/academy-go-q12021/model"
)

func FetchCryptocurreyncy() (string, error) {
	fiat := "MXN"
	crypto := "ETH"

	//Build the URL string
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat

	//Make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("an error occurred during the request, please try again")
	}

	defer resp.Body.Close()
	//Create a variable as the same type as our model
	var cResp model.Cryptocurrency

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("an error occurred while decoding, please try again")
	}
	return cResp.TextOuput(), nil

}
