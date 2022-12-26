// User data access object
// We will have entire logic to persist and retierve User from database.
package users

import (
	"fmt"

	"github.com/sharif-42/BookStore-User-Apis/data_sources/mysql/users_db"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/mysql_utils"
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

	stmt, err := users_db.Client.Prepare(QueryGetUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	// results, err :=stmt.Query(user.ID) we can use this as well but then we have to close the connection by
	// defer results.close() otherwise it will hold the database connection and we have run out of connections
	// As we only need to get a single row not multiple rows here so QueryRow() is perfect in this case.

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date, &user.Status); getErr != nil {
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

	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created_Date, user.Status, user.Password)
	fmt.Println("#########ERRRRRRRRRR", SaveErr)
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

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.ID)

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

func (user *User) FindByStatus(status string) (Users, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(QueryFindUserByStatus)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	// defer will only be executed when it executes a return, So if we put defer before the error then it may raise nil pointer error
	defer rows.Close() // as it returns rows and create an open connection, so we have to close this connection

	results := make([]User, 0) // as we don't know how many users is ther with the given status. so making a sloce with 0.
	// filling the slice
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created_Date, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return results, nil
}
