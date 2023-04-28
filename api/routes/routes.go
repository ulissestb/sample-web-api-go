package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func Routes(router *gin.Engine) {
	router.GET("/", helloWorld)
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, HelloWorld{Message: "hello, world!"})
}
