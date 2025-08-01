package config

import (
	"context"
	"fmt"
	"log"

	"github.com/BigWaffleMonster/Eventure_backend/api"
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events_abstractions"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type NewServerParams struct {
	fx.In

	AuthController *v1.AuthController
	CategoryController *v1.CategoryController
	EventController *v1.EventController
	ParticipantController *v1.ParticipantController
	UserController *v1.UserController
	DomainEventQueue domain_events_abstractions.DomainEventQueue
	ServerConfig utils.ServerConfig
}

func NewServer(lc fx.Lifecycle, p NewServerParams) {

	router := gin.Default()

	api.SwaggerInfo(p.ServerConfig)

	BuildPublicRoutes(router, p)

	BuildProtectedRoutes(router, p)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return OnStart(router, p)
		},
		OnStop: func(ctx context.Context) error {
			return OnStop()
		},
	})
 }

 func OnStart(router *gin.Engine, p NewServerParams) error{
	fmt.Println("Server starting...")

	go RunServer(router, p.ServerConfig)

	p.DomainEventQueue.StartQueue()
	
	return nil
 }

 func OnStop() error {
	fmt.Println("Server stopped")

	return nil
 }

 func RunServer(router *gin.Engine, config utils.ServerConfig) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	if err := router.Run(fmt.Sprintf(":%d", config.APP_PORT)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Server is running on port %d...\n", config.APP_PORT)
 }