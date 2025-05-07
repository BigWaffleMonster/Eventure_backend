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

func BuildServerConfig() ServerConfig{
	viper.SetConfigType("env")
	viper.SetConfigFile(CONFIG)

	if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("Error reading config file, %s", err)
    }

	return ServerConfig{
		DB_HOST: viper.GetString("DB_HOST"),
		DB_PORT: viper.GetInt("DB_PORT"),
		DB_USER: viper.GetString("DB_USER"),
		DB_PASSWORD: viper.GetString("DB_PASSWORD"),
		DB_NAME: viper.GetString("DB_NAME"),
		APP_PORT: viper.GetInt("APP_PORT"),
		JWT_SECRET: viper.GetString("JWT_SECRET"),
		JWT_SECRET_REFRESH: viper.GetString("JWT_SECRET_REFRESH"),
		GO_ENVIRONMENT: viper.GetString("GO_ENVIRONMENT"),
	}
}