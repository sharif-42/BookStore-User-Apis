// User data access object
// We will have entire logic to persist and retierve User from database.
package users

import (
	"fmt"

	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
	"github.com/sharif-42/BookStore-User-Apis/utils/time_utils"
)

var (
	// We will use MySQL later, for now we are using memory
	UserDB = make(map[int64]User)
)

func (user *User) Get() *errors.RestError {
	// Instead of creating function we are creating method here

	result, is_exists := UserDB[user.ID]
	// result will be actual value/data what we are looking for
	// and is exists return the boolean, That is is user exists with this key or not.

	if is_exists == false {
		return errors.NotFoundError(fmt.Sprintf("User %d Not Found!", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.Created_Date = result.Created_Date

	return nil
}

func (user *User) Save() *errors.RestError {
	// Instead of creating function we are creating method here

	current_user, is_exists := UserDB[user.ID]
	if is_exists {
		if current_user.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("Email %s Already Registered!", user.Email))
		}

		return errors.BadRequestError(fmt.Sprintf("User %d Already Exists!", user.ID))
	}

	user.Created_Date = time_utils.GetLocalNowTimeString()

	UserDB[user.ID] = *user
	return nil
}
