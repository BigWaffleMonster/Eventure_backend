package migrations

import (
	"fmt"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config utils.ServerConfig) (*gorm.DB, error) {
	retries := 5
	delay := 5 * time.Second

	for retries > 0 {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.DB_HOST,
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_NAME,
			config.DB_PORT,)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database!")
		}

		m := gormigrate.New(db, gormigrate.DefaultOptions, GetAllMigrations())
		err = m.Migrate()

		if err == nil {
			log.Println("Successfully migrate!")
			return db, nil
		}

		log.Printf("Failed to connect to database: %v. Retrying in %v... (%d attempts left)", err, delay, retries)
		retries--
		time.Sleep(delay)
	}

	log.Fatalf("Failed to connect to the database after multiple retries.")
	return nil, fmt.Errorf("failed to connect to the database after multiple retries")
}

func GetAllMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		M250820252255_Initial(),
		M250820252258_AddDomainEvents(),
		M030920252053_UpdateRefreshToken(),
		M030920252053_RemoveRefreshTokenFromSession(),
		// Добавляй сюда новые миграции
	}
}