package middleware

import (
	"errors"
	api "nagato/apiservice/pkg/plugin"

	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	myJwt "github.com/dokidokikoi/go-common/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		claims, err := myJwt.VerifyToken(tokenString, "test")

		if err != nil {
			ctx.Abort()
			if errors.Is(err, myErrors.ErrTokenExpired) {
				core.WriteResponse(ctx, myErrors.ApiErrTokenExpired, nil)
				return
			}
			core.WriteResponse(ctx, myErrors.ApiErrTokenValidation, nil)
			return
		}

		user, err := api.GetUserByEmail(ctx, claims.Emial, nil)
		if err != nil {
			ctx.Abort()
			return
		}

		ctx.Set("current_user", user)
		ctx.Next()
	}
}
