package repository

import (
	"context"
	"database/sql"
	"sample-web-api/internal/infra/db/models"
	db "sample-web-api/sql"
)

type ExampleRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewExampleRepository(dbt *sql.DB) *ExampleRepository {
	return &ExampleRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *ExampleRepository) InsertNewMessage(context context.Context, example *models.Example) error {
	_, err := r.Queries.InsertMessage(context, example.Message)

	if err != nil {
		return err
	}

	return nil
}

func (r *ExampleRepository) ListMessages(context context.Context) ([]models.Example, error) {
	examples, err := r.Queries.ListMessages(context)
	if err != nil {
		return nil, err
	}
	var exp []models.Example

	for _, e := range examples {
		exp = append(exp, *&models.Example{Message: e.Message})
	}

	return exp, nil
}

func (r *ExampleRepository) GetMessage(context context.Context, id int64) (models.Example, error) {
	example, err := r.Queries.GetMessage(context, id)
	if err != nil {
		return models.Example{}, err
	}
	return models.Example{Message: example.Message}, nil
}
