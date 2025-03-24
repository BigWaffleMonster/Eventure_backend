package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	retries := 5 // Maximum number of retries
	delay := 5 * time.Second

	for retries > 0 {
		// Build the DSN (Data Source Name) from environment variables
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))

		// Attempt to connect to the database
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database!")
		}

		err = db.SetupJoinTable(&models.User{}, "Events", &models.Participant{})
		db.AutoMigrate(&models.User{}, &models.Event{}, &models.Category{}, &models.Participant{}, &models.Notification{})

		if err == nil {
			log.Println("Successfully migrate!")
			return db, nil
		}

		// Log the error and retry after a delay
		log.Printf("Failed to connect to database: %v. Retrying in %v... (%d attempts left)", err, delay, retries)
		retries--
		time.Sleep(delay)
	}

	// If all retries fail, exit the program with an error
	log.Fatalf("Failed to connect to the database after multiple retries.")
	return nil, fmt.Errorf("failed to connect to the database after multiple retries")
}
