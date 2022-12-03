package domain

import "github.com/golang-jwt/jwt/v4"

type JwtClaim struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}
