package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int64  `json:"age"`
}

var Students = []Student{
	Student{ID: 1, Fullname: "Diego", Age: 31},
	Student{ID: 2, Fullname: "Mayara", Age: 28},
}

func hearthHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "healthy API",
	})
}

func listStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": Students,
	})
}

func getRoutes(gin *gin.Engine) *gin.Engine {
	gin.GET("/heart", hearthHandler)
	studentsGroupV1 := gin.Group("/api/v1")
	{
		studentsGroupV1.GET("/student", listStudentsHandler)
	}

	return gin
}

func main() {
	r := gin.Default()
	getRoutes(r)
	r.Run(":3333")
}
