// User data access object
// We will have entire logic to persist and retierve User from database.
package users

import (
	"fmt"
	"strings"

	"github.com/sharif-42/BookStore-User-Apis/data_sources/mysql/users_db"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/time_utils"
)

const (
	IndexUniqueEmail = "email_UNIQUE"
	ErrorNoRpws      = "no rows in result set"
	QueryInsertUser  = "INSERT INTO users(first_name, last_name, email, created_date) VALUES(?, ?, ?, ?);"
	QueGetUser       = "SELECT id, first_name, last_name, email, created_date FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestError {
	// Instead of creating function we are creating method here

	stmt, err := users_db.Client.Prepare(QueGetUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date); err != nil {
		// return errors.NotFoundError(fmt.Sprintf("User %d Not Found!", user.ID))
		if strings.Contains(err.Error(), ErrorNoRpws) {
			return errors.NotFoundError(fmt.Sprintf("User %d Not Found!", user.ID))
		}
		return errors.InternalServerError(fmt.Sprintf("Error while trying to get user %d: %s", user.ID, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	// Instead of creating function we are creating method here

	stmt, err := users_db.Client.Prepare(QueryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close() // After creating a statement we have to close it, otherwise it will hold the connection to database.
	// it will be executed just before the return statement of the function/method

	user.Created_Date = time_utils.GetLocalNowTimeString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created_Date)
	if err != nil {
		if strings.Contains(err.Error(), IndexUniqueEmail) {
			return errors.BadRequestError(fmt.Sprintf("Email %s Already Registered!", user.Email))
		}
		return errors.InternalServerError(fmt.Sprintf("Error while Saving the User: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("Error while Saving the User: %s", err.Error()))
	}

	// we can execute the query directly like the following
	// result, err = users_db.Client.Exec(QueryInsertUser, user.FirstName, user.LastName, user.Email, user.Created_Date)
	// But most of the people said that statement approach is better than the direct execute for most of the cases. and performance is better as well
	user.ID = userID

	return nil
}
