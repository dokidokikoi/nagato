package controller

import (
	"nagato/apiservice/internal/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c Controller) GetCurrentUser(ctx *gin.Context) *model.User {
	return ctx.MustGet("current_user").(*model.User)
}
