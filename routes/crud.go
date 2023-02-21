package routes

import (
	"zocket-assignment/controllers"

	"github.com/gin-gonic/gin"
)

func Crud(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)
	r.GET("/book/:id", controllers.GetBook)
	r.POST("/book", controllers.CreateBook)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
}
