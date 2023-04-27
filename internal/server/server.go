package server

import "github.com/gin-gonic/gin"

func Start() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello, world!",
		})
	})

	server.Run(":8080")
}
