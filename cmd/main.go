package main

import (
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/api"
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Eventura app
//	@version		1.0
//	@description	Simple app to plan your celebration.

//	@contact.name   Daniil, Sergei, Alex
//	@contact.email  rachkov.work@gmail.com, Sergei.m.khanlarov@gmail.com, me@justwalsdi.ru

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func main() {
	//TODO: Сепарировать данный метод
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print(err)
		log.Fatal("Error loading .env file")
		return
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userController := v1.NewUserController(userService)

	authService := auth.NewAuthService(userRepo)
	authController := v1.NewAuthController(authService)

	eventRepository := event.NewEventRepository(db)
	eventService := event.NewEventService(eventRepository)
	eventController := v1.NewEventController(eventService)

	participantRepository := participant.NewParticipantRepository(db)
	participantService := participant.NewParticipantService(participantRepository)
	participantController := v1.NewParticipantController(participantService)

	router := gin.Default()

	api.SwaggerInfo()

	public := router.Group("/api/v1")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.POST("/refresh", authController.RefreshToken)

		events := public.Group("/event")
		{
			events.POST("", eventController.Create)
			events.PUT("/:id", eventController.Update)
			events.DELETE("/:id", eventController.Remove)
			events.GET("/:id", eventController.GetByID)
			events.GET("", eventController.GetCollection)
		}
		participants := public.Group("/participant")
		{
			participants.POST("", participantController.Create)
			participants.PUT("/:id/state", participantController.ChangeState)
			participants.DELETE("/:id", participantController.Remove)
			participants.GET("/:id", participantController.GetByID)
			participants.GET("", participantController.GetCollection)
		}
	}

	protected := router.Group("/api/v1")
	{
		protected.Use(auth.AuthMiddleware())
		{
			user := public.Group("/user")
			{
				user.GET("/:id", userController.GetUserByID)
				user.PUT("/:id", userController.Update)
				user.DELETE("/:id", userController.Remove)
			}
		}
	}

	log.Println("Server is running on port 8080...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
