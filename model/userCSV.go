package model

type UserCSV struct {
	ID          int    `csv:"id"`
	Gender      string `csv:"gender"`
	Title       string `csv:"title"`
	First       string `csv:"first"`
	Last        string `csv:"last"`
	Email       string `csv:"email"`
	CellPhone   string `csv:"cell"`
	Nationality string `csv:"nat"`
}
