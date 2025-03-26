package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/notification"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
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

		err = db.SetupJoinTable(&event.Event{}, "Users", &participant.Participant{})
		db.AutoMigrate(&user.User{}, &event.Event{}, &category.Category{}, &participant.Participant{}, &notification.Notification{})

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
