package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func createStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	var student Student
	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not get payload",
		})
		return
	}
	student.ID = Students[len(Students)-1].ID + 1
	Students = append(Students, student)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": student,
	})
}

func updateStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("param: %d (type: %s) is required", id, "router parameter"),
		})
		return
	}

	var student Student
	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not get payload",
		})
		return
	}

	if student.Fullname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fullname is required",
		})
		return
	}

	if student.Age <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "age is required",
		})
		return
	}

	var studentFound *Student
	for _, s := range Students {
		if s.ID == id {
			studentFound = &s
			break
		}
	}

	if studentFound == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "student with id: " + strconv.Itoa(id) + " not found",
		})
		return
	}

	studentFound.Fullname = student.Fullname
	studentFound.Age = student.Age

	for i, s := range Students {
		if studentFound.ID == s.ID {
			Students[i].Fullname = studentFound.Fullname
			Students[i].Age = studentFound.Age
			break
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": studentFound,
	})

}

func deleteStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("param: %d (type: %s) is required", id, "router parameter"),
		})
		return
	}

	var student *Student
	indexFound := -1
	for i, s := range Students {
		if s.ID == id {
			student = &s
			indexFound = i
			break
		}
	}

	if indexFound == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "student with id: " + strconv.Itoa(id) + " not found",
		})
		return
	}

	Students = append(Students[:indexFound], Students[indexFound+1:]...)

	ctx.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func getRoutes(gin *gin.Engine) *gin.Engine {
	gin.GET("/heart", hearthHandler)
	studentsGroupV1 := gin.Group("/api/v1")
	{
		studentsGroupV1.GET("/student", listStudentsHandler)
		studentsGroupV1.POST("/student", createStudentsHandler)
		studentsGroupV1.PUT("/student/:id", updateStudentsHandler)
		studentsGroupV1.DELETE("/student/:id", deleteStudentsHandler)
	}
	return gin
}

func main() {
	r := gin.Default()
	getRoutes(r)
	r.Run(":3333")
}
