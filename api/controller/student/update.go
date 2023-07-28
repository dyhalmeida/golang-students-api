package student

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/dyhalmeida/golang-students-api/entity"
	"github.com/dyhalmeida/golang-students-api/entity/shared"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	var input StudentInput
	paramId := ctx.Param("id")
	UUID, err := shared.GetIDByString(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(
			fmt.Sprintf("param: %s (type: %s) is required or invalid uuid", paramId, "router parameter"),
		))
		return
	}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("could not get payload"))
		return
	}

	input.ID = UUID

	if input.Fullname == "" {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("fullname is required"))
		return
	}

	if input.Age <= 0 {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("age is required"))
		return
	}

	var studentFound *entity.Student
	for _, s := range entity.Students {
		if s.ID == input.ID {
			studentFound = &s
			break
		}
	}

	if studentFound == nil {
		ctx.JSON(http.StatusNotFound, controller.NewResponseMessageError(fmt.Sprintf("student with id: %s not found", paramId)))
		return
	}

	for i, s := range entity.Students {
		if input.ID == s.ID {
			entity.Students[i].Fullname = input.Fullname
			entity.Students[i].Age = input.Age
			break
		}
	}

	ctx.JSON(http.StatusOK, input)
}
