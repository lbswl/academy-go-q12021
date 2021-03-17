package model

import "fmt"

type Cryptocurrency struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

//TextOuput is exported, it formats the data to plain text
func (c Cryptocurrency) TextOuput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice : %s\nRank: %s\nHigh: %s\nCirculatingSupply: %s\n",
		c.Name, c.Price, c.Rank, c.High, c.CirculatingSupply)

	return p
}

func (c Cryptocurrency) TextOutputCSV() []string {

	cryptoLine := []string{c.Name, c.Price, c.Rank, c.High, c.CirculatingSupply}

	return cryptoLine
}
