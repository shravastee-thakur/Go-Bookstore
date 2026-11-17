package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	db *gorm.DB
)



func Connect() {
	// 1. Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// 2. Define your Postgres DSN (Connection String)
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN environment variable is required")
	}

	// 3. Open connection using the postgres driver
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db = d

}

// GetDB allows other packages to access the database instance
func GetDB() *gorm.DB  {
	return db
}