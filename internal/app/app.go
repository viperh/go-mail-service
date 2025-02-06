package app

import (
	"go-mail-service/internal/api/controllers"
	"go-mail-service/internal/api/routes"
	"go-mail-service/internal/pkg/config"
	"go-mail-service/internal/pkg/mail"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	API         *gin.Engine
	Config      *config.Config
	MailService *mail.MailService
	Controller  *controllers.Controller
}

func NewApp() *App {
	g := gin.Default()
	cfg := config.NewConfig()

	mailService := mail.NewMailService(cfg)

	control := controllers.NewController(mailService)

	routes.SetRoute(g, control)

	return &App{
		API:         g,
		Config:      cfg,
		MailService: mailService,
		Controller:  control,
	}

}

func (a *App) Run() {
	gin.SetMode(gin.DebugMode)
	err := a.API.Run(":" + a.Config.APIPort)
	if err != nil {
		log.Fatal(err)
	}
}
