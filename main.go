package main

import (
	"github.com/joho/godotenv"
	"github.com/kimoscloud/user-management-service/app/usecase"
	"github.com/kimoscloud/user-management-service/infrastructure/db"
	"github.com/kimoscloud/user-management-service/infrastructure/repository/postgres"
	"os"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".env.dev")
		if err != nil {
			panic("Error loading .env.dev file")
		}
	}
	connection, err := db.NewConnection()
	if err != nil {
		panic("error connecting to the database")
	}
	userRepository := postgres.NewUserRepository(connection)
	_ = usecase.NewCreateUserUseCase(userRepository)

}
