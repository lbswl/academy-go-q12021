package model

type UserJSON struct {
	ID          int    `json:"id"`
	Gender      string `json:"gender"`
	Title       string `json:"title"`
	First       string `json:"first"`
	Last        string `json:"last"`
	Email       string `json:"email"`
	CellPhone   string `json:"cell"`
	Nationality string `json:"nat"`
}
