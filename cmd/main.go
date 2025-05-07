package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/api"
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/BigWaffleMonster/Eventure_backend/api/middlewares"
)

//	@title			Eventura app
//	@version		1.0
//	@description	Simple app to plan your celebration.

//	@contact.name   Daniil, Sergei, Alex
//	@contact.email  rachkov.work@gmail.com, Sergei.m.khanlarov@gmail.com, me@justwalsdi.ru

//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	//TODO: Сепарировать данный метод
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Print(err)
	// 	log.Fatal("Error loading .env file")
	// 	return
	// }

	// db, err := utils.InitDB()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	// userRepo := user.NewUserRepository(db)
	// userService := user.NewUserService(userRepo)
	// userController := v1.NewUserController(userService)

	// authService := auth.NewAuthService(userRepo)
	// authController := v1.NewAuthController(authService)

	// eventRepository := event.NewEventRepository(db)
	// eventService := event.NewEventService(eventRepository)
	// eventController := v1.NewEventController(eventService)

	// participantRepository := participant.NewParticipantRepository(db)
	// participantService := participant.NewParticipantService(participantRepository)
	// participantController := v1.NewParticipantController(participantService)

	// router := gin.Default()

	// api.SwaggerInfo()

	// public := router.Group("/api/v1")
	// public.Use(middlewares.HandleCors())
	// {
	// 	public.POST("/register", authController.Register)
	// 	public.POST("/login", authController.Login)
	// 	public.POST("/refresh", authController.RefreshToken)

	// 	events := public.Group("/event")
	// 	{
	// 		events.GET("/:id", eventController.GetByID)
	// 		events.GET("", eventController.GetCollection)
	// 	}
	// }

	// protected := router.Group("/api/v1")
	// protected.Use(middlewares.HandleCors(), middlewares.AuthMiddleware())
	// {
	// 	{
	// 		user := public.Group("/user")
	// 		{
	// 			user.GET("/:id", userController.GetByID)
	// 			user.PUT("/:id", userController.Update)
	// 			user.DELETE("/:id", userController.Remove)
	// 		}
	// 		events := public.Group("/event")
	// 		{
	// 			events.POST("", eventController.Create)
	// 			events.PUT("/:id", eventController.Update)
	// 			events.DELETE("/:id", eventController.Remove)
	// 		}
	// 		participants := public.Group("/participant")
	// 		{
	// 			participants.POST("", participantController.Create)
	// 			participants.PUT("/:id/state", participantController.ChangeState)
	// 			participants.DELETE("/:id", participantController.Remove)
	// 			participants.GET("/:id", participantController.GetByID)
	// 			participants.GET("", participantController.GetCollection)
	// 		}

	// 	}
	// }

	// log.Println("Server is running on port 8080...")
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	// if err := router.Run(":8080"); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }


	app := fx.New(
		event.AddDI(),
		user.AddDI(),
		participant.AddDI(),
		category.AddDI(),
		fx.Provide(provideGormDB),
		fx.Provide(auth.NewAuthService),
		fx.Provide(v1.NewAuthController),
		fx.Provide(v1.NewEventController),
		fx.Provide(v1.NewCategoryController),
		fx.Provide(v1.NewUserController),
		fx.Provide(v1.NewParticipantController),
		fx.Invoke(NewServer),
	)

	app.Run()
}

func provideGormDB() (*gorm.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print(err)
		log.Fatal("Error loading .env file")
		return nil, err
	}

	return utils.InitDB()
}

type NewServerParams struct {
	fx.In

	AuthController v1.AuthController
	CategoryController v1.CategoryController
	EventController v1.EventController
	ParticipantController v1.ParticipantController
	UserController v1.UserController
}

func NewServer(lc fx.Lifecycle, p NewServerParams) {

	router := gin.Default()

	api.SwaggerInfo()

	public := router.Group("/api/v1")
	public.Use(middlewares.HandleCors())
	{
		public.POST("/register", p.AuthController.Register)
		public.POST("/login", p.AuthController.Login)
		public.POST("/refresh", p.AuthController.RefreshToken)

		events := public.Group("/event")
		{
			events.GET("/:id", p.EventController.GetByID)
			events.GET("", p.EventController.GetCollection)
		}
		categories := public.Group("/category")
		{
			categories.GET("/:id", p.CategoryController.GetByID)
			categories.GET("", p.CategoryController.GetCollection)
		}
	}

	protected := router.Group("/api/v1")
	protected.Use(middlewares.HandleCors(), middlewares.AuthMiddleware())
	{
		{
			user := public.Group("/user")
			{
				user.GET("/:id", p.UserController.GetByID)
				user.PUT("/:id", p.UserController.Update)
				user.DELETE("/:id", p.UserController.Remove)
			}
			events := public.Group("/event")
			{
				events.POST("", p.EventController.Create)
				events.PUT("/:id", p.EventController.Update)
				events.DELETE("/:id", p.EventController.Remove)
			}
			participants := public.Group("/participant")
			{
				participants.POST("", p.ParticipantController.Create)
				participants.PUT("/:id/state", p.ParticipantController.ChangeState)
				participants.DELETE("/:id", p.ParticipantController.Remove)
				participants.GET("/:id", p.ParticipantController.GetByID)
				participants.GET("", p.ParticipantController.GetCollection)
			}

		}
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return OnStart()
		},
		OnStop: func(ctx context.Context) error {
			return OnStop()
		},
	})

	// if err := router.Run(":8080"); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
	log.Println("Server is running on port 8080...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	go router.Run(":5500")
 }

 func OnStart() error{
	fmt.Println("Server started")
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Print(err)
	// 	log.Fatal("Error loading .env file")
	// 	return err
	// }

	// _ , err := utils.InitDB()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	return nil
 }

 func OnStop() error {
	fmt.Println("Server stopped")

	return nil
 }