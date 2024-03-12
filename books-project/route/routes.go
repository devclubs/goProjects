package route

import "github.com/gin-gonic/gin"

func CollectRoute(r *gin.Engine) *gin.Engine {

	registerUserRouter(r)
	registerBookRouter(r)
	return r
}
