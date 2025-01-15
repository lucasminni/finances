package db

import (
	"finances/internal/domain/models/debt"
	"finances/internal/domain/models/income"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SQLConnector *gorm.DB

func ConnectDatabase() {

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

	SQLConnector, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: " + err.Error())
	} else {
		log.Println("Connected to database!")
	}

	if os.Getenv("AUTOMIGRATE") == "true" {
		log.Println("Starting AutoMigrate...")
		err = SQLConnector.AutoMigrate(
			&debt.Debt{},
			&income.Income{},
		)
		if err != nil {
			log.Fatalf("Failed to auto migrate: %v", err)
		} else {
			log.Println("Migration complete!")
		}
	}

	if os.Getenv("DEBUG") == "true" {
		SQLConnector.Debug()
		log.Println("Debug mode on!")
	}
	
}

func CheckDatabase() bool {
	db, _ := SQLConnector.DB()
	err := db.Ping()

	if err != nil {
		return false
	}
	return true
}
