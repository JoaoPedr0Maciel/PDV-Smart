package routes

import (
	"pdv/handlers"
	"pdv/handlers/company"
	"pdv/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("v1")
	{
		v1.GET("hello-world", handlers.HelloWorld)
		v1.POST("users", user.CreateUserHandler)


		v1.POST("companies", company.CreateCompanyHandler)
	}
}
