package services

import (
	"github.com/studingprojects/bookstore_user-api/domain/users"
	"github.com/studingprojects/bookstore_user-api/utils/crypto_utils"
	"github.com/studingprojects/bookstore_user-api/utils/date_utils"
	"github.com/studingprojects/bookstore_user-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindByStatus(string) (users.Users, *errors.RestErr)
	Login(string, string) (*users.User, *errors.RestErr)
}

type usersService struct {
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowString()
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(id int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: id}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	current, getErr := s.GetUser(userId)
	if getErr != nil {
		return getErr
	}
	if deleteErr := current.Delete(); deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (s *usersService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) Login(email string, password string) (*users.User, *errors.RestErr) {
	user := &users.User{Email: email, Password: crypto_utils.GetMd5(password)}
	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
