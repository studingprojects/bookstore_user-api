package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	errors "github.com/studingprojects/bookstore_utils-go/rest_errors"
)

const (
	NoRecordPattern = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NoRecordPattern) {
			return errors.NewNotFounfError("record not found")
		}
		return errors.NewInternalServerError("error while parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request", err)
}
