// Map All urls from here
package app

import (
	"github.com/sharif-42/BookStore-User-Apis/controllers/ping"
	"github.com/sharif-42/BookStore-User-Apis/controllers/users"
)

func MapUrls() {
	router.GET("ping/", ping.Ping)

	// User Apis
	router.GET("users/:user_id", users.GetUser)
	router.GET("users/search", users.SearchUser)

	router.POST("users/", users.CreateUser)

}
