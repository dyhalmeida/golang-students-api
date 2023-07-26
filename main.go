package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hearthHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy API",
	})
}

func getRoutes(gin *gin.Engine) *gin.Engine {
	gin.GET("/heart", hearthHandler)
	return gin
}

func main() {
	r := gin.Default()
	getRoutes(r)
	r.Run(":3333")
}
