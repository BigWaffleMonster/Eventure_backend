package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port           string
	Mode           string
	AllowedOrigins []string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️  Не удалось прочитать .env: %v (использую переменные окружения)", err)
		return nil, err
	}

	var config = Config{
		Server: ServerConfig{
			Port:           getEnv("PORT", "8080"),
			Mode:           getEnv("GIN_MODE", "debug"),
			AllowedOrigins: viper.GetStringSlice("SERVER_ALLOWED_ORIGINS"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "eventure_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	fmt.Print(config.Server.AllowedOrigins, "JEHFWEFHI#@*)*IOUJKH")

	if len(config.Server.AllowedOrigins) == 0 {
		config.Server.AllowedOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	}

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	if val := viper.GetString(key); val != "" {
		return val
	}
	return defaultValue
}
