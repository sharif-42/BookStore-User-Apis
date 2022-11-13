// User Data Transfer Object. We are transferring data to persistant layer that means
// from Database to application and vice-verca.

package users

import (
	"strings"

	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Created_Date string `json:"created_date"`
}

// func Validate(user *User) *errors.RestError {
// 	// This is function, which is taking user as input and validating the user
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.BadRequestError("Invalid Email address!")
// 	}
// 	return nil
// }

func (user *User) Validate() *errors.RestError {
	// This is a method of user struct not an individual function
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid Email address!")
	}
	return nil
}
