package client

type UserSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
