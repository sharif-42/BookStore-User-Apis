package services

import (
	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
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

	if err := user.Validate(); err != nil {
		// there is error while validating data
		return nil, err
	}

	if err := user.Save(); err != nil {
		// there is error while creating user
		return nil, err
	}
	// no error so return the user.
	return &user, nil
}
