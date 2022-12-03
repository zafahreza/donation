package client

type UserGetNewOtpRequest struct {
	Email string `validate:"required,email" json:"email"`
}
