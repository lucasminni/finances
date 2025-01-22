package db

import (
	"finances/internal/domain/models/debt"
	"finances/internal/domain/models/income"
	"finances/internal/infra/settings"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var SQLConnector *gorm.DB

func ConnectDatabase() {
	var err error

	host := settings.GetEnvs().DatabaseHost
	user := settings.GetEnvs().DatabaseUser
	password := settings.GetEnvs().DatabasePassword
	name := settings.GetEnvs().DatabaseName
	port := settings.GetEnvs().DatabasePort
	ssl := settings.GetEnvs().DatabaseSSL
	automigrate := settings.GetEnvs().DatabaseAutoMigrate
	debug := settings.GetEnvs().DatabaseDebug

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		name,
		port,
		ssl)

	SQLConnector, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: " + err.Error())
	} else {
		log.Println("Connected to database!")
	}

	if automigrate == "true" {
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

	if debug == "true" {
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
