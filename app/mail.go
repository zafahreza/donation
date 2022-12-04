package app

import (
	"donation/helper"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
	"os"
	"time"
)

func NewSmtpClient() *mail.SMTPClient {
	err := godotenv.Load("mail.env")
	helper.PanicIfError(err)

	username := os.Getenv("SMTP_NAME")
	password := os.Getenv("SMTP_PASS")

	smtp := mail.NewSMTPClient()
	smtp.Host = "smtp.gmail.com"
	smtp.Port = 587
	smtp.Username = username
	smtp.Password = password
	smtp.Encryption = mail.EncryptionTLS

	smtp.SendTimeout = 10 * time.Second
	smtp.ConnectTimeout = 10 * time.Second
	smtp.KeepAlive = true

	smtpClient, err := smtp.Connect()
	helper.PanicIfError(err)

	return smtpClient
}
