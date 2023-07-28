package student

import (
	"net/http"

	"github.com/dyhalmeida/golang-students-api/entity"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, entity.Students)
}
