package middleware

import (
	"time"

	"github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
)

type requestParams struct {
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		start := time.Now()

		ctx.Next()

		param := requestParams{}
		param.Latency = time.Since(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path
		zap.L().Sugar().Infof("接口调用		%v", param)
	}
}
