package student

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/dyhalmeida/golang-students-api/entity"
	"github.com/dyhalmeida/golang-students-api/entity/shared"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	paramId := ctx.Param("id")

	UUID, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(
			fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
		))
		return
	}

	indexStudentFound := -1
	var student *entity.Student

	for i, s := range entity.Students {
		if s.ID == UUID {
			student = &s
			indexStudentFound = i
			break
		}
	}

	if indexStudentFound == -1 {
		ctx.JSON(http.StatusNotFound, controller.NewResponseMessageError(fmt.Sprintf("student with id: %s not found", UUID)))
		return
	}

	entity.Students = append(entity.Students[:indexStudentFound], entity.Students[indexStudentFound+1:]...)

	ctx.JSON(http.StatusOK, student)
}
