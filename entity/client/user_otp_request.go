package client

type UserOtpRequest struct {
	Email string `validate:"required,email" json:"email"`
	OTP   string `validate:"required,len=5" json:"otp"`
}
