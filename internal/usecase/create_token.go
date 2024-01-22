package usecase

import (
	"bank-backend/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *domain.User, secretKey string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		Issuer: user.GetLogin(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}