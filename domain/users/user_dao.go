package users

import (
	"fmt"
	"strings"

	"github.com/studingprojects/bookstore_utils-go/logger"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"

	"github.com/studingprojects/bookstore_user-api/datasources/userdb"
	"github.com/studingprojects/bookstore_user-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, password, status, date_created) VALUES (?, ?, ?, ?, ?, ?);"
	queryFindByID               = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE email=? AND password=?;"
)

func (user *User) Get() rest_errors.RestErr {
	stmt, err := userdb.UserDb.Prepare(queryFindByID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() rest_errors.RestErr {
	stmt, clientErr := userdb.UserDb.Prepare(queryInsertUser)
	if clientErr != nil {
		return clientErr
	}
	defer stmt.Close()

	saveResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := saveResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
			err,
		)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := userdb.UserDb.Prepare(queryUpdateUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), nil)
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := userdb.UserDb.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), nil)
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, clientErr := userdb.UserDb.Prepare(queryFindByStatus)
	if clientErr != nil {
		logger.Error("error when try to prepare get user statement", clientErr)
		return nil, clientErr
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, clientErr := userdb.UserDb.Prepare(queryFindByEmailAndPassword)
	if clientErr != nil {
		logger.Error("error when try to prepare get user statement", clientErr)
		return clientErr
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysql_utils.NoRecordPattern) {
			return rest_errors.NewBadRequestError("invalid user credentials")
		}
		logger.Error("Error when trying to get user by email & password", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	return nil
}
