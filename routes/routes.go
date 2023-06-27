package routes

import (
	"book-gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", controller.GetBooks)
	router.POST("/", controller.CreateBook)
	router.DELETE("/:id", controller.DeleteBook)
	router.PUT("/:id", controller.UpdateBook)
}
