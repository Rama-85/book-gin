package main

import (
	"book-gin/config"
	"book-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
