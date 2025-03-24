package main

import (
	"log"

	docs "github.com/BigWaffleMonster/Eventure_backend/cmd/docs"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/services"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/infrastructure/repositories"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/infrastructure/utils"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/presentation/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   Daniil, Sergei, Alex
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	eventRepo := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	// Настройка маршрутизатора
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Eventura app"
	docs.SwaggerInfo.Description = "Simple app to plan your celebration"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5001"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api := router.Group("/api/v1")
	{
		api.POST("/register", authController.Register)
		api.POST("/event", eventController.Create)
	}

	// Запуск сервера
	log.Println("Server is running on port 5001...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := router.Run(":5001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
