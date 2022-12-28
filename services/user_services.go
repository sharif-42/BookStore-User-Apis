package services

import (
	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/utils/crypto_utils"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/time_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
	// creating a variable of UsersService of type usersServiceInterface and creating new instance of usersService
)

type usersService struct {
	// an emty struct for user service. so that all methods belong to this service
}

type usersServiceInterface interface {
	// this interface will help us in many ways like followings
	// 1. We can mock, while we run the tests, otherwise mock will not possible
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(users.User, bool) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
}

func (usersService *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.BadRequestError("Invalid User Id!")
	}
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (usersService *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	// validate the data before save in to database
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

func (usersService *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {
	// update the user and return updated user or error if there is any
	current, err := usersService.GetUser(user.ID)
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

func (usersService *usersService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{ID: userId}

	// checking is user really exists by given ID
	_, getError := usersService.GetUser(userId)
	if getError != nil {
		return getError
	}

	// User exists and now we can perform delete
	return user.Delete()
}

func (usersService *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	userObj := &users.User{}
	return userObj.FindByStatus(status)
}
