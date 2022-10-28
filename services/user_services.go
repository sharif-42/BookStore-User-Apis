package services

import (
	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
