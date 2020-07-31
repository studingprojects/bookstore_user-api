package model

type User struct {
	BaseModel
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
	Password    string `json:"-"`
}
