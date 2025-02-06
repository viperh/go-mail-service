package routes

import (
	"go-mail-service/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoute(g *gin.Engine, cnt *controllers.Controller) {
	g.POST("/sendMail", cnt.SendMail)
}
