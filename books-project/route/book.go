package route

import (
	"books_project/controller"

	"github.com/gin-gonic/gin"
)

func registerBookRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/book/add", controller.CreateBookHandler)
		v1.GET("/book/list", controller.GetBookListHandler)
		v1.GET("/book/detial/:id", controller.GetBookDetailHandler)
		v1.PUT("/book/update", controller.UpdateBookDetailHandler)
	}
}
