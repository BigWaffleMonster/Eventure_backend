package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/api"
	v1 "github.com/BigWaffleMonster/Eventure_backend/api/v1"
	"github.com/BigWaffleMonster/Eventure_backend/config"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	DomainEventQueue interfaces.DomainEventQueue
	Logger sglogger.Logger
}

func NewServer(lc fx.Lifecycle, p NewServerParams) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	
	router := gin.Default()
	router.Use(gin.Recovery())

	api.SwaggerInfo()

    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "OK",
            "message": "Service is running",
        })
    })

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

    router.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/health")
    })
	
	router.GET("/swagger", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })

	BuildRoutes(router, p)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			p.Logger.Info(ctx, "Starting application components...")
			go RunServer(ctx, router, p)

			p.DomainEventQueue.StartQueue(ctx)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Info(ctx, "Shutting down application components...")
			return nil
		},
	})
 }

 func RunServer(ctx context.Context, router *gin.Engine, p NewServerParams) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	if err := router.Run(fmt.Sprintf(":%d", config.GetAppPort())); err != nil {
		p.Logger.Fatal(ctx, "Failed to start server: %s", err.Error())
		log.Fatalf("Failed to start server: %v", err)
	}

	p.Logger.Info(ctx, "Server started successfully on port %d", config.GetAppPort())
 }