package main

import (
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/controller"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/service/auth_service"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print(err)
		log.Fatal("Error loading .env file")
		return
	}
	// Подключение к PostgreSQL
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автомиграция

	// Инициализация слоев
	userRepo := repository.NewUserRepository(db)
	authService := auth_service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	// Настройка маршрутизатора
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		// api.POST("/login", authController.Login)
	}

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
