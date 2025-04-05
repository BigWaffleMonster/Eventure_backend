package main

import (
	"log"

	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
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
	userService := user.NewUserService(userRepo)
	userController := v1.NewUserController(userService)

	authService := auth.NewAuthService(userRepo)
	authController := v1.NewAuthController(authService)

	// Настройка маршрутизатора
	router := gin.Default()

	// controllers := []any{*authController}
	// routes.SetupRoutes(router, controllers)

	public := router.Group("/api/v1")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		// public.POST("/refresh-token", authController.RefreshToken)
	}

	protected := router.Group("/api/v1")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/user/:id", userController.GetUserByID)
	}

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
