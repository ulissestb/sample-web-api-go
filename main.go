package main

import (
	"database/sql"
	"fmt"
	"sample-web-api/api/routes"
	"sample-web-api/api/server"
	"sample-web-api/configs"
	"sample-web-api/internal/domain"
	"sample-web-api/internal/infra/db/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := configs.LoadConfigs(".")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	fmt.Println(dataSource)
	db, err := sql.Open(config.DBDriver, dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exampleRepository := repository.NewExampleRepository(db)
	useCase := domain.NewHelloWorldUseCase(exampleRepository)
	srv := server.Start()
	routes.Routes(srv, &useCase)
	srv.Run(":" + config.AppPort)
}
