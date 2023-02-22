package middleware

import (
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/dokidokikoi/go-common/core"
	myErrors "github.com/dokidokikoi/go-common/errors"
	"github.com/gin-gonic/gin"
)

func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("harukaze"), nil
		})

		if err != nil {
			ctx.Abort()
			if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
				core.WriteResponse(ctx, myErrors.ApiErrTokenExpired, nil)
				return
			}
			core.WriteResponse(ctx, myErrors.ApiErrTokenValidation, nil)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		store, err := data.GetStoreDBFactory()
		if err != nil {
			ctx.Abort()
			return
		}
		user, err := store.Users().Get(ctx, &model.User{Email: claims["sub"].(string)}, nil)
		if err != nil {
			ctx.Abort()
			return
		}

		ctx.Set("current_user", user)
		ctx.Next()
	}
}
