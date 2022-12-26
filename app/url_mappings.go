// Map All urls from here
package app

import (
	"github.com/sharif-42/BookStore-User-Apis/controllers/ping"
	"github.com/sharif-42/BookStore-User-Apis/controllers/users"
)

func MapUrls() {
	router.GET("ping/", ping.Ping)

	// User Apis
	router.POST("users/", users.CreateUser)
	router.GET("users/:user_id", users.GetUser)
	router.PUT("users/:user_id", users.UpdateUser)
	router.PATCH("users/:user_id", users.UpdateUser)
	router.DELETE("users/:user_id", users.DeleteUser)
	router.GET("internal/users/search", users.SearchUserByStatus)

}
