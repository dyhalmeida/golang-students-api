package main

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-students-api/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Age      int64     `json:"age"`
}

var Students = []Student{
	{ID: shared.GetID(), Fullname: "Diego", Age: 31},
	{ID: shared.GetID(), Fullname: "Mayara", Age: 28},
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

func showStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	paramId := ctx.Param("id")
	id, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
		})
		return
	}

	var student *Student
	for _, s := range Students {
		if s.ID == id {
			student = &s
			break
		}
	}

	if student == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Errorf("student with id: %s not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": student,
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
	student.ID = shared.GetID()
	Students = append(Students, student)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": student,
	})
}

func updateStudentsHandler(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	paramId := ctx.Param("id")
	id, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
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
			"message": fmt.Sprintf("student with id: %s not found", paramId),
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

	paramId := ctx.Param("id")
	id, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
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
			"message": fmt.Errorf("student with id: %s not found", id),
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
		studentsGroupV1.GET("/student/:id", showStudentsHandler)
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
