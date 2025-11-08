package migrations

import (
	"context"
	"fmt"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/config"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(logger sglogger.Logger) *gorm.DB {
	ctx := context.Background()
	retries := 5
	delay := 5 * time.Second

	for retries > 0 {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.GetDBHost(),
			config.GetDBUser(),
			config.GetDBPassword(),
			config.GetDBName(),
			config.GetDBPort(),)

		tmp := viper.AllKeys()
		fmt.Print(tmp)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			logger.Info(ctx, "Successfully connected to the database!")
		}

		m := gormigrate.New(db, gormigrate.DefaultOptions, GetAllMigrations())
		err = m.Migrate()

		if err == nil {
			logger.Info(ctx, "Successfully migrate!")
			return db
		}

		logger.Error(ctx, "Failed to connect to database: %v. Retrying in %v... (%d attempts left)", err, delay, retries)

		retries--
		time.Sleep(delay)
	}

	logger.Fatal(ctx, "Failed to connect to the database after multiple retries.")

	return nil
}

func GetAllMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		M250820252255_Initial(),
		M250820252258_AddDomainEvents(),
		M030920252053_UpdateRefreshToken(),
		M030920252053_RemoveRefreshTokenFromSession(),
		//TODO: Добавляй сюда новые миграции
	}
}