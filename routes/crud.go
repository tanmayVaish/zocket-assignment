package routes

import (
	"zocket-assignment/controllers"

	"github.com/gin-gonic/gin"
)

func CrudRoute(r *gin.Engine) {

	crud := r.Group("/crud")

	crud.GET("/books", controllers.GetBooks)
	crud.GET("/book/:id", controllers.GetBook)
	crud.POST("/book", controllers.CreateBook)
	crud.PUT("/book/:id", controllers.UpdateBook)
	crud.DELETE("/book/:id", controllers.DeleteBook)
}
