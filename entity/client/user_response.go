package client

import "time"

type UserResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	UpdatedAt time.Time `json:"updated_at"`
}
