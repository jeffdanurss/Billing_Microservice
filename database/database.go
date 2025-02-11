package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "fmt"
    "github.com/joho/godotenv"
    "billing-microservice/models"
)

var DB *gorm.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Failed to load .env file")
    }

    // Connect initializes the database conection
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    log.Println("Connected to the database!")
     // Migrate
    err = DB.AutoMigrate(&models.Invoice{}) 
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
    log.Println("Database migration completed!")
}