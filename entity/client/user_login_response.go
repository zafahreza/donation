package client

import "time"

type UserLoginResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Token     string    `json:"token"`
	UpdatedAt time.Time `json:"updated_at"`
}
