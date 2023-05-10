package inittask

import (
	"nagato/apiservice/internal/config"
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/heartbeat"
	"nagato/apiservice/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	config.Init("./internal/conf/application.yml")

	go heartbeat.ListenHeartbeat()
	data.SetStoreDBFactory()
	// client.InitClients()
	initIndex()

	r.UseRawPath = true
	router.InitRoutes(r)
}
