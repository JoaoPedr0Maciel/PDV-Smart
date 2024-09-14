package routes

import (
	"pdv/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("v1")
	{
		v1.GET("", handlers.HelloWorld)
	}
}
