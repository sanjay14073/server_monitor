package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
	"github.com/joho/godotenv"
)

func main() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//get url
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("URL environment variable is not set")
	}
	//get credentials
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	maintainUser:=os.Getenv("MAINTAIN_USER")
	if smtpUser == "" || smtpPassword == "" {
		log.Fatal("SMTP credentials or maintaince user email are not set in .env file")
	}

	d := gomail.NewDialer("smtp.gmail.com", 587, smtpUser, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc("*/30 * * * * *", func() {
		if _, err := http.Get(url); err != nil {
			message := gomail.NewMessage()
			message.SetHeader("From", smtpUser)
			message.SetHeader("To",maintainUser)
			message.SetHeader("Subject", "Web Service Down")
			message.SetBody("text/plain", "Your web service is down, please check it.")
	
			err := d.DialAndSend(message)
			if err != nil {
				log.Println("Failed to send email:", err)
			}
		}
	})

	if err != nil {
		log.Fatal("Error adding cron job:", err)
	}

	//Start CRON
	c.Start()

	//Terminate cron
	defer c.Stop()

	select {}
}
