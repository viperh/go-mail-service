package app

import (
	"go-mail-service/internal/api/controllers"
	"go-mail-service/internal/api/middlewares"
	"go-mail-service/internal/api/routes"
	"go-mail-service/internal/pkg/config"
	"go-mail-service/internal/pkg/mail"
	"go-mail-service/internal/pkg/rabbit"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	API           *gin.Engine
	Config        *config.Config
	MailService   *mail.MailService
	JwtMiddleware *middlewares.JwtMiddleware
	Controller    *controllers.Controller
	Consumer      *rabbit.Consumer
}

func NewApp() *App {
	var g *gin.Engine
	var jwtMiddleware *middlewares.JwtMiddleware
	var controller *controllers.Controller
	var consumer *rabbit.Consumer

	cfg := config.NewConfig()
	mailService := mail.NewMailService(cfg)
	if cfg.Mode == "api" {
		g = gin.Default()

		jwtMiddleware := middlewares.NewMiddleware(cfg)
		g.Use(jwtMiddleware.JwtMiddlewareFunc)
		controller := controllers.NewController(mailService)
		routes.SetRoute(g, controller)
	}

	if cfg.Mode == "rabbitmq" {
		consumer = rabbit.NewRabbitConsumer(cfg, mailService)
	}

	return &App{
		API:           g,
		Config:        cfg,
		MailService:   mailService,
		JwtMiddleware: jwtMiddleware,
		Controller:    controller,
		Consumer:      consumer,
	}

}

func (a *App) Run() {
	gin.SetMode(gin.ReleaseMode)
	err := a.API.Run(":" + a.Config.APIPort)
	if err != nil {
		log.Fatal(err)
	}
}
