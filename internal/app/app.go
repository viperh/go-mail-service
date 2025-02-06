package app

import (
	"go-mail-service/internal/api/controllers"
	"go-mail-service/internal/api/routes"
	"go-mail-service/internal/pkg/config"
	"go-mail-service/internal/pkg/mail"

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

	cntrl := controllers.NewController(mailService)

	routes.SetRoute(g, cntrl)

	return &App{
		API:         g,
		Config:      cfg,
		MailService: mailService,
		Controller:  cntrl,
	}

}

func (a *App) Run() {
	a.API.Run(a.Config.APIPort)
}
