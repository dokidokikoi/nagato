package user

import (
	"github.com/dokidokikoi/go-common/core"
	"github.com/gin-gonic/gin"
)

func (c UserController) Me(ctx *gin.Context) {
	u := c.Controller.GetCurrentUser(ctx)
	core.WriteResponse(ctx, nil, u)
}
