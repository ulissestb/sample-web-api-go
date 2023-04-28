package server

import (
	"sample-web-api/api/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()
	routes.Routes(server)
	server.Run(":8080")
}
