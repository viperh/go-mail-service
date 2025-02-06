package controllers

import (
	"go-mail-service/internal/api/controllers/dto"
	"go-mail-service/internal/pkg/mail"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	MailService *mail.MailService
}

func NewController(m *mail.MailService) *Controller {
	return &Controller{
		MailService: m,
	}
}

func (cn *Controller) SendMail(c *gin.Context) {

	var req dto.SendMailReq

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cn.MailService.SendMail(req.From, req.To, req.Subject, req.Body, req.IsHtml)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "What  the f"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mail sent successfully!"})

}
