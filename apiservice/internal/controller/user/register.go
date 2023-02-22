package user

import (
	"nagato/apiservice/internal/model"

	"github.com/dokidokikoi/go-common/core"
	"github.com/dokidokikoi/go-common/crypto"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func (c UserController) Register(ctx *gin.Context) {
	input := UserRegister{}
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	pwd, err := crypto.EncryptPassword(input.Password)
	if err != nil {
		zaplog.L().Error("用户密码加密错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}
	creatUser := &model.User{
		Email:    input.Email,
		Username: input.Username,
		Password: pwd,
		Avatar:   "",
	}

	if err := c.service.User().Create(ctx, creatUser); err != nil {
		zaplog.L().Error("数据库插入用户数据错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabaseOp, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("注册成功"), nil)
}
