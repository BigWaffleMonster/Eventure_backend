package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	c "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
)

func InitDB(config *c.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if config.Server.Mode == "debug" {
		if err := db.AutoMigrate(&schema.User{}, &schema.Category{}, &schema.Event{}, &schema.Participant{}); err != nil {
			log.Printf("⚠️  Ошибка миграции: %v", err)
		} else {
			log.Println("✅ Миграции выполнены")
		}
	}

	return db, nil
}
