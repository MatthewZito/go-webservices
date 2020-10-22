package app

import (
	"github.com/MatthewZito/go-microservices/mvc/controllers"
)

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}
