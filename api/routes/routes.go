package routes

import (
	"net/http"
	"sample-web-api/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func Routes(router *gin.Engine, usecase *domain.HelloWorldUseCase) {
	router.GET("/example", func(ctx *gin.Context) {
		usecase.ListMessages(ctx)
	})

	router.POST("/example", func(ctx *gin.Context) {
		var message HelloWorld
		if err := ctx.BindJSON(&message); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		usecase.AddNewMessage(ctx, message.Message)
	})

	router.GET("/example/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}

		usecase.FindOneMessage(ctx, int(id))
	})
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, HelloWorld{Message: "hello, world!"})
}
