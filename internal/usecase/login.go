package usecase

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type LoginCommand struct {
	Login string
	Password []byte
}

type LoginUseCase struct {
	userRepository domain.UserRepository
}

func NewLoginUseCase(userRepository domain.UserRepository) *LoginUseCase{
	return &LoginUseCase{
		userRepository: userRepository,
	}
}

func (useCase *LoginUseCase) LoginHandler(ctx context.Context, command *LoginCommand) (string, error) {
		user, err := useCase.userRepository.FindByName(ctx, command.Login)
		
		if err != nil {
			slog.Error("user not found", err)
			return "", err
		}

		err = bcrypt.CompareHashAndPassword(user.GetPasswordHash(), command.Password)

		if err != nil {
			slog.Error("password incorrect", err)
			return "", err
		}

		return CreateToken(user, os.Getenv("SECRET_KEY"))
}