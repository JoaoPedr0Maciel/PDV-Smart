package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"hello": "world",
	})
}
