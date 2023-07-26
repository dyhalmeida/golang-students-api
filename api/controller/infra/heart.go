package infra

import (
	"net/http"

	"github.com/dyhalmeida/golang-students-api/api/controller"
	"github.com/gin-gonic/gin"
)

func HeartHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.NewResponseMessage("healthy API"))
}
