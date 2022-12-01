package domain

import (
	"time"
)

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	Bio          string
	ImageUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
