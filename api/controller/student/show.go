package student

import (
	"fmt"
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/dyhalmeida/golang-students-api/entity/shared"
	student_usecase "github.com/dyhalmeida/golang-students-api/usecases/student"
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

	student, err := student_usecase.ShowStudentByID(UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, student)

}
