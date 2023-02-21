package controllers

import (
	"zocket-assignment/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	var books []models.Book
	models.DB.Find(&books)
	c.JSON(200, gin.H{
		"status": "success",
		"books":  books,
	})
}

func GetBook(c *gin.Context) {

	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Book not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"book":   book,
	})
}

func CreateBook(c *gin.Context) {

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid JSON",
		})
		return
	}

	models.DB.Create(&book)

	c.JSON(201, gin.H{
		"status": "success",
		"book":   book,
	})
}

func UpdateBook(c *gin.Context) {

	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Book not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid JSON",
		})
		return
	}

	models.DB.Save(&book)

	c.JSON(200, gin.H{
		"status": "success",
		"book":   book,
	})
}

func DeleteBook(c *gin.Context) {

	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Book not found",
		})
		return
	}

	models.DB.Delete(&book)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Book deleted",
	})
}
