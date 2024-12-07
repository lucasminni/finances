package db

import (
	"financas/internal/schemas/debt"
	"financas/internal/schemas/income"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v, err")
	} else {
		log.Println("Connected to database!")
	}

	fmt.Println("Starting AutoMigrate...")
	err = db.AutoMigrate(
		&debt.Debt{},
		&income.Income{},
	)

	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	} else {
		log.Println("Migration complete!")
	}

}
