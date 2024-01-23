package middlewares

import (
	"context"
	"github.com/google/uuid"
)

type key struct {
}

func GetUserIdFromContext(ctx context.Context) uuid.UUID {
	token, _ := ctx.Value(key{}).(uuid.UUID)

	return token
}