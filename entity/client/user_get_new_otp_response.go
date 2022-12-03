package client

type UserGetNewOtpResponse struct {
	Email string `json:"email"`
	Msg   string `json:"msg"`
}
