// User data access object
// We will have entire logic to persist and retierve User from database.
package users

import (
	"github.com/sharif-42/BookStore-User-Apis/data_sources/mysql/users_db"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/mysql_utils"
	"github.com/sharif-42/BookStore-User-Apis/utils/time_utils"
)

const (
	QueryInsertUser = "INSERT INTO users(first_name, last_name, email, created_date) VALUES(?, ?, ?, ?);"
	QueryGetUser    = "SELECT id, first_name, last_name, email, created_date FROM users WHERE id=?;"
	QueryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	QueryDeleteUser = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestError {
	// Instead of creating function we are creating method here

	stmt, err := users_db.Client.Prepare(QueryGetUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	// results, err :=stmt.Query(user.ID) we can use this as well but then we have to close the connection by
	// defer results.close() otherwise it will hold the database connection and we have run out of connections
	// As we only need to get a single row not multiple rows here so QueryRow() is perfect in this case.

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date); getErr != nil {
		return mysql_utils.ParseError(getErr)
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
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created_Date)
	if SaveErr != nil {
		return mysql_utils.ParseError(SaveErr)
	}

	userID, insertError := insertResult.LastInsertId()
	if insertError != nil {
		return mysql_utils.ParseError(insertError)
	}
	// we can execute the query directly like the following
	// result, err = users_db.Client.Exec(QueryInsertUser, user.FirstName, user.LastName, user.Email, user.Created_Date)
	// But most of the people said that statement approach is better than the direct execute for most of the cases. and performance is better as well

	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(QueryUpdateUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)

	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(QueryDeleteUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, deleteError := stmt.Exec(user.ID)

	if deleteError != nil {
		return mysql_utils.ParseError(deleteError)
	}
	return nil
}
