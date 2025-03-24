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

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	// Настройка маршрутизатора
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/event", eventController.Create)
	}

	// Запуск сервера
	log.Println("Server is running on port 5001...")
	if err := router.Run(":5001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
