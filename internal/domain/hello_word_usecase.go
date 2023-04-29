package domain

import (
	"net/http"
	"sample-web-api/internal/infra/db/models"
	"sample-web-api/internal/infra/db/repository"

	"github.com/gin-gonic/gin"
)

type HelloWorldUseCase struct {
	exampleRepository repository.ExampleRepository
}

func NewHelloWorldUseCase(er *repository.ExampleRepository) HelloWorldUseCase {
	return HelloWorldUseCase{
		exampleRepository: *er,
	}
}

func (r *HelloWorldUseCase) ListMessages(ctx *gin.Context) {
	examples, err := r.exampleRepository.ListMessages(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(422, err)
	}

	if examples == nil {
		ctx.AbortWithStatusJSON(422, gin.H{"message": "nenhum exemplo cadastrado"})
		return
	}

	ctx.JSON(http.StatusOK, examples)
}

func (r *HelloWorldUseCase) AddNewMessage(ctx *gin.Context, message string) {
	err := r.exampleRepository.InsertNewMessage(ctx.Request.Context(), models.NewExample(message))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}

func (r *HelloWorldUseCase) FindOneMessage(ctx *gin.Context, id int) {
	example, err := r.exampleRepository.GetMessage(ctx.Request.Context(), int64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, example)
}
