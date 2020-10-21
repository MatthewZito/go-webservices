package app

import (
	"log"
	"net/http"

	"github.com/MatthewZito/go-microservices/mvc/controllers"
)

const port = ":8080"

func InitApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(port, nil); err == nil {
		log.Panic(err)
	}

}
