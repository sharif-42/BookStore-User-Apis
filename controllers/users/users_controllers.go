package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharif-42/BookStore-User-Apis/domain/users"
	"github.com/sharif-42/BookStore-User-Apis/services"
	"github.com/sharif-42/BookStore-User-Apis/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

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
		rest_err := errors.BadRequestError("Invalid Json Body!")
		c.JSON(rest_err.Status, rest_err)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := errors.BadRequestError("Invalid User Id!")
		c.JSON(err.Status, err)
		return
	}
	fmt.Println(userId)
	result, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "User Search is not implemented yet!")
}
