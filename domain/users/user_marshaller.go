package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
