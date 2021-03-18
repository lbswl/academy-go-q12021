package model

type Response struct {
	Results []Results `json:"results,omitempty"`
	Info    Info      `json:"info,omitempty"`
}

type Results struct {
	Gender      string `json:"gender,omitempty"`
	Name        Name   `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Cell        string `json:"cell,omitempty"`
	Nationality string `json:"nat,omitempty"`
}

type Name struct {
	Title string `json:"title,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

type Info struct {
	Seed          string `json:"seed,omitempty"`
	NumberResults int    `json:"results,omitempty"`
	Page          int    `json:"page,omitempty"`
	Version       string `json:"version,omitempty"`
}
