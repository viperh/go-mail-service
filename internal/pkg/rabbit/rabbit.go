package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go-mail-service/internal/pkg/config"
	"go-mail-service/internal/pkg/mail"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type MailMessage struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	IsHtml  bool     `json:"is_html"`
}

type Consumer struct {
	conn        *amqp.Connection
	ch          *amqp.Channel
	q           amqp.Queue
	MailService *mail.MailService
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitConsumer(cfg *config.Config, mailService *mail.MailService) *Consumer {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		cfg.RabbitQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	defer conn.Close()
	defer ch.Close()

	return &Consumer{
		conn:        conn,
		ch:          ch,
		q:           q,
		MailService: mailService,
	}
}

func (c *Consumer) Consume() {
	msgs, err := c.ch.Consume(
		c.q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var mailMessage MailMessage
			err := json.Unmarshal(d.Body, &mailMessage)
			if err != nil {
				log.Println("Error decoding JSON")
				continue
			}
			err = c.MailService.SendMail(mailMessage.From, mailMessage.To, mailMessage.Subject, mailMessage.Body, mailMessage.IsHtml)
			if err != nil {
				log.Println("Error sending mail")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	log.Printf("Shutting down consumer")

}
