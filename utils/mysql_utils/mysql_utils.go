package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

const (
	IndexUniqueEmail = "email_UNIQUE"
	ErrorNoRows      = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, isSqlErr := err.(*mysql.MySQLError)
	// we are checking the given error is a mysql error or not
	if !isSqlErr {
		// that means this is not a mysql error
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NotFoundError("No record matching by given id!")
		}
		return errors.InternalServerError("Error parsing database response!")
	}
	// this error is a mysql error and we have to handle it
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("Invalid data!")
		// TODO: we will handle multiple numbers here
	}
	return errors.InternalServerError("error processing request!")
}
