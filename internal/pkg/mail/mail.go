package mail

import (
	"go-mail-service/internal/pkg/config"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMailService(cfg *config.Config) *MailService {
	return &MailService{
		Host:     cfg.SMTPTarget,
		Port:     cfg.SMTPPort,
		Username: cfg.SMTPUser,
		Password: cfg.SMTPPassword,
	}
}

func (m *MailService) SendMail(from string, to []string, subject string, body string, isHtml bool) error {
	message := gomail.NewMessage()

	message.SetHeader("From", from)
	length := len(to)
	message.SetHeader("To", to[:length]...)
	message.SetHeader("Subject", subject)
	if isHtml {
		message.SetBody("text/html", body)
	} else {
		message.SetBody("text/plain", body)
	}

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil

}
