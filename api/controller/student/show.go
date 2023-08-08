package student

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/dyhalmeida/golang-students-api/entity"
	"github.com/dyhalmeida/golang-students-api/entity/shared"
	"github.com/gin-gonic/gin"
)

func Show(ctx *gin.Context) {
	paramId := ctx.Param("id")
	UUID, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(
			fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
		))
		return
	}

	var student *entity.Student
	for _, s := range entity.Students {
		if s.ID == UUID {
			student = &s
			break
		}
	}

	if student == nil {
		ctx.JSON(http.StatusNotFound, controller.NewResponseMessageError(fmt.Sprintf("student with id: %s not found", UUID)))
		return
	}

	ctx.JSON(http.StatusOK, student)

}
