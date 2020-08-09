package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"
)

const (
	NoRecordPattern = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NoRecordPattern) {
			return rest_errors.NewNotFounfError("record not found")
		}
		return rest_errors.NewInternalServerError("error while parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request", err)
}
