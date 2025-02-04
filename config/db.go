package config

import (
    "fmt"
	"github.com/joho/godotenv"
	"os"
	"log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// setting database conection
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=5432 sslmode=disable", 
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
