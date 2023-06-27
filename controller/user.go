package controller

import (
	"book-gin/config"
	"book-gin/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)
	c.JSON(200, &books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	config.DB.Create(&book)
	c.BindJSON(&book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id=?", c.Param("id")).First(&book)
	c.BindJSON(&book)
	config.DB.Save(&book)
	c.JSON(200, &book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id=?", c.Param("id")).Delete(&book)
	c.JSON(200, &book)
}
