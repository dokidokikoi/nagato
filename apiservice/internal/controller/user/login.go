package user

import (
	"nagato/apiservice/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dokidokikoi/go-common/core"
	"github.com/dokidokikoi/go-common/crypto"
	myErrors "github.com/dokidokikoi/go-common/errors"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c UserController) Login(ctx *gin.Context) {
	input := &UserLogin{}
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	user, err := c.service.User().Get(ctx, &model.User{Email: input.Email}, &meta.GetOption{Include: []string{"email"}})
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrRecordNotFound, nil)
		return
	}

	if !crypto.CheckPassword(input.Password, user.Password) {
		core.WriteResponse(ctx, myErrors.ApiErrPassword, nil)
		return
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		zaplog.L().Error("获取jwt token失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{"token": token})
}

func GenerateToken(subject string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(60 * 60 * 24 * time.Second)
	issuer := "harukaze"
	claims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  nowTime.Unix(),
		NotBefore: nowTime.Unix(),
		Issuer:    issuer,
		Subject:   subject,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("harukaze"))
	return token, err
}
