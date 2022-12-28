// User data access object
// We will have entire logic to persist and retierve User from database.
package users

import (
	"fmt"

	"github.com/sharif-42/BookStore-User-Apis/data_sources/mysql/users_db"
	"github.com/sharif-42/BookStore-User-Apis/logger"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

const (
	QueryInsertUser       = "INSERT INTO users(first_name, last_name, email, created_date, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	QueryGetUser          = "SELECT id, first_name, last_name, email, created_date, status FROM users WHERE id=?;"
	QueryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	QueryDeleteUser       = "DELETE FROM users WHERE id=?;"
	QueryFindUserByStatus = "SELECT id, first_name, last_name, email, created_date, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestError {
	// Instead of creating function we are creating method here

	stmt, stmtErr := users_db.Client.Prepare(QueryGetUser)
	if stmtErr != nil {
		error_message := "Error while trying to prepare get user statement"
		logger.Error(error_message, stmtErr)
		return errors.InternalServerError("Database Error!")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	// results, err :=stmt.Query(user.ID) we can use this as well but then we have to close the connection by
	// defer results.close() otherwise it will hold the database connection and we have run out of connections
	// As we only need to get a single row not multiple rows here so QueryRow() is perfect in this case.

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date, &user.Status); getErr != nil {
		error_message := fmt.Sprintf("Error while trying to get user by ID=%d", user.ID)
		logger.Error(error_message, getErr)
		return errors.InternalServerError("Database Error!")
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	// Instead of creating function we are creating method here

	stmt, stmtErr := users_db.Client.Prepare(QueryInsertUser)
	if stmtErr != nil {
		logger.Error("Error while trying to prepare create user statement", stmtErr)
		return errors.InternalServerError("Database Error!")
	}
	defer stmt.Close() // After creating a statement we have to close it, otherwise it will hold the connection to database.
	// it will be executed just before the return statement of the function/method

	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created_Date, user.Status, user.Password)
	if SaveErr != nil {
		logger.Error("Error while trying to create user", SaveErr)
		return errors.InternalServerError("Database Error!")
	}

	userID, insertError := insertResult.LastInsertId()
	if insertError != nil {
		logger.Error("Error while trying to get last insert id after creating a new user", insertError)
		return errors.InternalServerError("Database Error!")
	}
	// we can execute the query directly like the following
	// result, err = users_db.Client.Exec(QueryInsertUser, user.FirstName, user.LastName, user.Email, user.Created_Date)
	// But most of the people said that statement approach is better than the direct execute for most of the cases. and performance is better as well

	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, stmtErr := users_db.Client.Prepare(QueryUpdateUser)
	if stmtErr != nil {
		logger.Error("Error while trying to prepare update user statement", stmtErr)
		return errors.InternalServerError("Database Error!")
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.ID)
	if updateErr != nil {
		logger.Error("Error while trying to update user", updateErr)
		return errors.InternalServerError("Database Error!")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, stmtErr := users_db.Client.Prepare(QueryDeleteUser)
	if stmtErr != nil {
		logger.Error("Error while trying to prepare delete user statement", stmtErr)
		return errors.InternalServerError("Database Error!")
	}
	defer stmt.Close()

	_, deleteError := stmt.Exec(user.ID)
	if deleteError != nil {
		logger.Error("Error while trying to delete user", deleteError)
		return errors.InternalServerError("Database Error!")
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestError) {
	stmt, stmtErr := users_db.Client.Prepare(QueryFindUserByStatus)
	if stmtErr != nil {
		logger.Error("Error while trying to prepare find user by status statement", stmtErr)
		return nil, errors.InternalServerError("Database Error!")
	}
	defer stmt.Close()

	rows, findErr := stmt.Query(status)
	if findErr != nil {
		logger.Error("Error while trying to delete user", findErr)
		return nil, errors.InternalServerError("Database Error!")
	}
	// defer will only be executed when it executes a return, So if we put defer before the error then it may raise nil pointer error
	defer rows.Close() // as it returns rows and create an open connection, so we have to close this connection

	results := make([]User, 0) // as we don't know how many users is ther with the given status. so making a sloce with 0.
	// filling the slice
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date, &user.Status); err != nil {
			logger.Error("Error while scan user row into user struct", err)
			return nil, errors.InternalServerError("Database Error!")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return results, nil
}
