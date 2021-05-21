package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - Returns pointer to a new database struct
func NewDatabase() (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	table := os.Getenv("DB_TABLE")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, table, password)
	fmt.Println(connectionString)

	database, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return database, err
	}

	if err := database.DB().Ping(); err != nil {
		return database, err
	}
	return database, nil
}
