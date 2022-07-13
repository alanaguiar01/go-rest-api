package database

import (
	"fmt"
	"os"

	"github.com/alanaguiar01/go-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	// variables localized in .env
	postgresUser := os.Getenv("POSTGRES_DB")
	postgresDb := os.Getenv("POSTGRES_DB")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	//config gorm with database
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432", postgresUser, postgresPassword, postgresDb)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//verify connection database
	if err != nil {
		return nil, err
	}
	//make auto migrate database
	db.AutoMigrate(&models.User{})
	return db, nil
}
