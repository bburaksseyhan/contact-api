package model

type Contact struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
