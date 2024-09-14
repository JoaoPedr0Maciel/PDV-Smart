package routes

import "github.com/gin-gonic/gin"

func InitRouter() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run()
}
