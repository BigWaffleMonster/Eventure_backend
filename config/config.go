package config

import (
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()

	viper.SetConfigFile(".env")
	viper.MergeInConfig()

	viper.AutomaticEnv()
}

func GetConsoleLoggingLevel() string {
    return getStringWithDefault("logging__console__level", "info")
}

func GetLokiLoggingLevel() string {
    return getStringWithDefault("logging__loki__level", "info")
}

func GetAppName() string {
    return getStringWithDefault("app__name", "eventura")
}

func GetLokiUrl() string {
    return viper.GetString("logging__loki__url")
}

func GetDBHost() string {
	return getStringWithDefault("db_host", "localhost")
}

func GetDBPort() int {
	return getIntWithDefault("db__port", 5432)
}

func GetDBUser() string {
	return getStringWithDefault("db__user", "postgres")
}

func GetDBPassword() string {
	return getStringWithDefault("db__password", "")
}

func GetDBName() string {
	return getStringWithDefault("db__name", "app")
}

// Application
func GetAppPort() int {
	return getIntWithDefault("app__port", 8080)
}

// JWT
func GetJWTSecret() string {
	return getStringWithDefault("jwt__secret", "default-jwt-secret-change-in-production")
}

func GetJWTRefreshSecret() string {
	return getStringWithDefault("jwt__secret_refresh", "default-refresh-secret-change-in-production")
}

// Environment
func GetGoEnvironment() string {
	return getStringWithDefault("go_environment", "development")
}

func GetRequestTimeout() int {
	return getIntWithDefault("request_timeout", 30)
}

// Helper functions
func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func getStringWithDefault(key string, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntWithDefault(key string, defaultValue int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultValue
}