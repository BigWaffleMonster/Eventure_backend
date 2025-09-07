package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var CONFIG = "../.env"

type ServerConfig struct{
	DB_HOST string
	DB_PORT int
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	APP_PORT int

	JWT_SECRET string
	JWT_SECRET_REFRESH string

	GO_ENVIRONMENT string
}

func BuildServerConfig() ServerConfig {
	viper.SetConfigType("env")
	viper.SetConfigFile(CONFIG)

	// Пытаемся прочитать из файла, но не падаем если его нет
	err := viper.ReadInConfig();

	if err != nil {
		fmt.Printf("Note: Config file not found, using environment variables: %s\n", err)

		// Автоматически читает переменные окружения с префиксом
		viper.AutomaticEnv()
		
		// Устанавливаем привязки для переменных окружения
		viper.BindEnv("DB_HOST")
		viper.BindEnv("DB_PORT")
		viper.BindEnv("DB_USER")
		viper.BindEnv("DB_PASSWORD")
		viper.BindEnv("DB_NAME")
		viper.BindEnv("APP_PORT")
		viper.BindEnv("JWT_SECRET")
		viper.BindEnv("JWT_SECRET_REFRESH")
		viper.BindEnv("GO_ENVIRONMENT")
	}

	return ServerConfig{
		DB_HOST:            viper.GetString("DB_HOST"),
		DB_PORT:            viper.GetInt("DB_PORT"),
		DB_USER:            viper.GetString("DB_USER"),
		DB_PASSWORD:        viper.GetString("DB_PASSWORD"),
		DB_NAME:            viper.GetString("DB_NAME"),
		APP_PORT:           viper.GetInt("APP_PORT"),
		JWT_SECRET:         viper.GetString("JWT_SECRET"),
		JWT_SECRET_REFRESH: viper.GetString("JWT_SECRET_REFRESH"),
		GO_ENVIRONMENT:     viper.GetString("GO_ENVIRONMENT"),
	}
}