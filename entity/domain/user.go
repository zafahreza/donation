package domain

import "time"

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	ImageUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
