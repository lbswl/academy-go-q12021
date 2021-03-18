package model

type UserCSV struct {
	Id          int    `csv:"id"`
	Gender      string `csv:"gender"`
	Title       string `csv:"title"`
	First       string `csv:"first"`
	Last        string `csv:"last"`
	Email       string `csv:"email"`
	CellPhone   string `csv:"cell"`
	Nationality string `csv:"nat"`
}
