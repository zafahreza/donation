package middleware

import (
	"donation/entity/domain"
	"donation/exception"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthMiddlewareImpl struct {
}

func NewAuthMiddleware() AuthMidleware {
	return &AuthMiddlewareImpl{}
}
func loadSecretKey() []byte {
	err := godotenv.Load("jwt.env")
	if err != nil {
		log.Fatal("gagal load env")
	}
	//helper.PanicIfError(err)
	secret := os.Getenv("SECRET_KEY")

	byteSecret := []byte(secret)
	return byteSecret
}

func (authMiddleware *AuthMiddlewareImpl) GenerateToken(userId int) string {

	secret := loadSecretKey()

	expDate := time.Now().Add(5 * time.Minute)

	claims := &domain.JwtClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("gagal generate token")
	}
	//helper.PanicIfError(err)

	return signedToken

}

func (authMiddleware *AuthMiddlewareImpl) ValidateToken(r *http.Request) int {
	header := r.Header.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		panic(exception.NewUnauthorizedError(errors.New("unauthorized, not contain bearer")))
	}

	arrayHeader := strings.Split(header, " ")
	tokenString := arrayHeader[1]

	claims := &domain.JwtClaim{}

	secret := loadSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claim, ok := token.Claims.(*domain.JwtClaim); ok && token.Valid {
		userId := claim.UserId

		return userId
	}

	jwtError, _ := err.(*jwt.ValidationError)
	if jwtError.Errors == jwt.ValidationErrorExpired {
		panic(exception.NewUnauthorizedError(errors.New("token expired, please login")))
	}
	if jwtError.Errors == jwt.ValidationErrorSignatureInvalid {
		panic(exception.NewUnauthorizedError(errors.New("unauthorized, invalid signature")))
	}

	panic(exception.NewUnauthorizedError(errors.New("unauthorized, token invalid")))
}
