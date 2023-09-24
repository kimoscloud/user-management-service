package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func NewConnection() (*gorm.DB, error) {
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASS")
	databaseHost := os.Getenv("DB_HOST")
	databasePort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Error parsing the database port")
	}
	databaseName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseHost, databasePort, databaseUser, databasePassword, databaseName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
