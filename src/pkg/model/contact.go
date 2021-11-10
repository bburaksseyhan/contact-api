package model

type Contact struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
