package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

var router *gin.Engine

func init() {
	router = gin.Default() // vs .New() - recovers from panic
}

func InitApp() {
	mapURLs()

	if err := router.Run(port); err != nil {
		log.Panic(err)
	}

}
