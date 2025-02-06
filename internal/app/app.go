package app

import (
	"go-mail-service/internal/api/controllers"
	"go-mail-service/internal/api/middlewares"
	"go-mail-service/internal/api/routes"
	"go-mail-service/internal/pkg/config"
	"go-mail-service/internal/pkg/mail"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	API           *gin.Engine
	Config        *config.Config
	MailService   *mail.MailService
	JwtMiddleware *middlewares.JwtMiddleware
	Controller    *controllers.Controller
}

func NewApp() *App {
	g := gin.Default()

	cfg := config.NewConfig()

	mailService := mail.NewMailService(cfg)

	middleware := middlewares.NewMiddleware(cfg)

	g.Use(middleware.JwtMiddlewareFunc)

	control := controllers.NewController(mailService)

	routes.SetRoute(g, control)

	return &App{
		API:           g,
		Config:        cfg,
		MailService:   mailService,
		JwtMiddleware: middleware,
		Controller:    control,
	}

}

func (a *App) Run() {
	gin.SetMode(gin.ReleaseMode)
	err := a.API.Run(":" + a.Config.APIPort)
	if err != nil {
		log.Fatal(err)
	}
}
