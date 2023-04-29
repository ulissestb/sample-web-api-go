package server

import (
	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())

	return server
}
