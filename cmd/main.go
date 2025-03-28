package main

import (
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/api"
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	userRepo := user.NewUserRepository(db)
	authService := user.NewAuthService(userRepo)
	authController := v1.NewAuthController(authService)

	eventRepository := event.NewEventRepository(db)
	eventService := event.NewEventService(eventRepository)
	eventController := v1.NewEventController(eventService)

	// Настройка маршрутизатора
	router := gin.Default()

	api.SwaggerInfo()

	public := router.Group("/api")
	{
		v1 := public.Group("/v1")
		v1.POST("/register", authController.Register)
		v1.POST("/login", authController.Login)
		v1.POST("/event", eventController.Create)
		v1.PUT("/event/:id", eventController.Update)
		v1.DELETE("/event/:id", eventController.Delete)
		v1.GET("/event/:id", eventController.GetById)
		v1.GET("/event", eventController.GetCollection)
		// public.POST("/refresh-token", authController.RefreshToken)
	}

	protected := router.Group("/api")
	{
		//v1 := public.Group("/v1")
	}
	protected.Use(auth.AuthMiddleware())

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
