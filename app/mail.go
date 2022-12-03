package app

import (
	"donation/helper.go"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
	"os"
)

func NewSmtpClient() *mail.SMTPClient {
	err := godotenv.Load("mail.env")
	helper.PanicIfError(err)
	//host := os.Getenv("SMTP_HOST")
	username := os.Getenv("SMTP_NAME")
	password := os.Getenv("SMTP_PASS")
	//port := os.Getenv("SMTP_PORT")
	//
	//smtpPort, err := strconv.Atoi(port)
	//helper.PanicIfError(err)

	smtp := mail.NewSMTPClient()
	smtp.Host = "smtp.gmail.com"
	smtp.Port = 587
	smtp.Username = username
	smtp.Password = password
	smtp.Encryption = mail.EncryptionTLS

	smtpClient, err := smtp.Connect()
	helper.PanicIfError(err)

	return smtpClient
}
