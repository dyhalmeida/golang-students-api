package student

import (
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/dyhalmeida/golang-students-api/entity"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var input StudentInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("could not get payload"))
		return
	}
	student := entity.NewStudent(input.Fullname, input.Age)
	entity.Students = append(entity.Students, *student)
	ctx.JSON(http.StatusOK, student)
}
