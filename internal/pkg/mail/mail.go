package mail

import "go-mail-service/internal/pkg/config"

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

func (m *MailService) SendMail(from string, to string, subject string, body string) error {
	return nil
}
