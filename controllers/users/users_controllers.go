package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/services"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

func getUserByID(userIdParam string) (int64, *errors.RestError) {
	userId, UserErr := strconv.ParseInt(userIdParam, 10, 64)
	if UserErr != nil {
		return 0, errors.BadRequestError("User Id should be a number!")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: We have to handle the error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO: We have to handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid Json Body!")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	// Validating user id
	userId, idErr := getUserByID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User

	// validating requested data
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid Json Body!")
		c.JSON(restErr.Status, restErr)
		return
	}

	// updating the user by given requested data
	user.ID = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateError := services.UpdateUser(user, isPartial)
	if updateError != nil {
		c.JSON(updateError.Status, updateError)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func GetUser(c *gin.Context) {
	userId, idErr := getUserByID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	result, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserByID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if deleteError := services.DeleteUser(userId); deleteError != nil {
		c.JSON(deleteError.Status, deleteError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchUserByStatus(c *gin.Context) {
	status := c.Query("status") // reading query param from the URL
	users, searchErr := services.SearchUser(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}
	users.Marshall(c.GetHeader("X-Public") == "true")

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
