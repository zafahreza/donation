package helper

import (
	"crypto/rand"
	"donation/entity/domain"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"io"
)

func GenerateOtp() string {
	max := 5
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func SendOtp(otp domain.OTP, smtp *mail.SMTPClient) {

	body := "Please input Your Verification Code: " + otp.OTP + " to complete registration"

	email := mail.NewMSG()
	email.SetFrom("Donation Website <fahrezaspam@gmail.com>")
	email.AddTo(otp.Email)
	email.SetSubject("Verification Code")
	email.SetBody(mail.TextPlain, body)
	err := email.Send(smtp)
	PanicIfError(err)

	fmt.Println("email sent to user")
}
