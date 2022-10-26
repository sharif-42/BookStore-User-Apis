// Map All urls from here
package app

import (
	"github.com/sharif-42/BookStore-User-Apis/controllers"
)

func MapUrls() {
	router.GET("ping/", controllers.Ping)

}
