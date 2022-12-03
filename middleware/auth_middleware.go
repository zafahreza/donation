package middleware

import "net/http"

type AuthMidleware interface {
	GenerateToken(userId int) string
	ValidateToken(r *http.Request) int
}
