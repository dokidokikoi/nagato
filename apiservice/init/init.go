package inittask

import (
	"nagato/apiservice/internal/config"
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"
	"nagato/apiservice/rpc/client"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	config.Init("./internal/conf/application.yml")

	go heartbeat.ListenHeartbeat()
	data.SetStoreDBFactory()
	client.InitClients()

	// 不将base64编码的路径解码
	r.UseRawPath = true
	router.InitRoutes(r)
}
