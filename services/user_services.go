package services

import (
	"fmt"

	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/utils/crypto_utils"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/time_utils"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.BadRequestError("Invalid User Id!")
	}
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	// validate the data before save in to database
	fmt.Println(user)
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Created_Date = time_utils.GetNowDBFormat()
	user.Status = users.StatusPending // For newly created user, status will be pending
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {
	// update the user and return updated user or error if there is any
	current, err := GetUser(user.ID)
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
		if user.Status != "" {
			current.Status = user.Status
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{ID: userId}

	// checking is user really exists by given ID
	_, getError := GetUser(userId)
	if getError != nil {
		return getError
	}

	// User exists and now we can perform delete
	return user.Delete()
}

func SearchUser(status string) (users.Users, *errors.RestError) {
	userObj := &users.User{}
	return userObj.FindByStatus(status)
}
