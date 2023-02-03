package main

import (
	"nagato/gateway/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InstallRoutes(r)

	r.Run(":8900")
}
