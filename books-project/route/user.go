package route

import (
	"books_project/controller"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(r *gin.Engine) {

	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
}
