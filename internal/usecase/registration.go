package usecase

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationCommand struct {
	Login string
	Password []byte
}

type RegistrationUseCase struct {
	userRepository domain.UserRepository
}

func NewRegistrationUseCase(userRepository domain.UserRepository) *RegistrationUseCase {
	return &RegistrationUseCase{
		userRepository: userRepository,
	}
}


func (useCase *RegistrationUseCase) RegistrationHandler(ctx context.Context, command *RegistrationCommand) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(command.Password, bcrypt.DefaultCost)

	if err != nil {
		slog.Error("generating password", err)
		return "", err
	}

	user := domain.NewUser(uuid.New(), command.Login, passwordHash)
	err = useCase.userRepository.Save(ctx, user)

	if err != nil {
		slog.Error("save to database", err)
		return "", err
	}

	return CreateToken(user, os.Getenv("SECRET_KEY"))
}
