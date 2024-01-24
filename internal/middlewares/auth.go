package middlewares

import (
	"context"

	"github.com/google/uuid"
)

type key struct {
}

type Claims struct {

}

func GetUserIdFromContext(ctx context.Context) uuid.UUID {
	token, _ := ctx.Value(key{}).(uuid.UUID)

	return token
}

//var errUnauthorized = errors.New("unauthorized")

// func Auth(writer http.ResponseWriter, request *http.Request) {
// 	authHeader := request.Header.Get("Authorization")

// 	token, err := jwt.ParseWithClaims(authHeader, &Claims{}, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Metho
// 	})
// }