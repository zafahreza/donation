package helper

import (
	"bytes"
	"crypto/rand"
	"donation/entity/domain"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	template2 "html/template"
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

func template(otp domain.OTP) bytes.Buffer {
	var body bytes.Buffer
	t := template2.Must(template2.ParseFiles("helper/template/template.html"))
	t.Execute(&body, domain.OTP{
		OTP: otp.OTP,
	})

	return body
}

func SendOtp(otp domain.OTP, smtp *mail.SMTPClient) {

	body := template(otp)

	email := mail.NewMSG()
	email.SetFrom("Donation Website <fahrezaspam@gmail.com>")
	email.AddTo(otp.Email)
	email.SetSubject("Verification Code")
	email.SetBody(mail.TextHTML, string(body.Bytes()))
	err := email.Send(smtp)
	PanicIfError(err)

	fmt.Println("email sent to user")
}
