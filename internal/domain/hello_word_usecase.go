package domain

import (
	"sample-web-api/internal/infra/db/repository"
)

type HelloWorldUseCase struct {
	exampleRepository repository.ExampleRepository
}

func NewHelloWorldUseCase(er *repository.ExampleRepository) HelloWorldUseCase {
	return HelloWorldUseCase{
		exampleRepository: *er,
	}
}

// func (er *HelloWorldUseCase) GetOneExample(id int) models.Example {
// 	return er.exampleRepository.GetExample(id)
// }

// func (er *HelloWorldUseCase) ListExamples() []models.Example {

// }
