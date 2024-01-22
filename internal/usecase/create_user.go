package usecase

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateCommand struct {
	Login string
	Email string
	Password []byte
}

type CreateUserUseCase struct {
	userRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}


func (useCase *CreateUserUseCase) CreateUserHandler(ctx context.Context, command *CreateCommand) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(command.Password, bcrypt.DefaultCost)

	if (err != nil) {
		slog.Error("generating password", err)
		return "", err
	}

	user := domain.NewUser(uuid.New(), command.Login, command.Email, passwordHash)
	err = useCase.userRepository.Save(ctx, user)

	if err != nil {
		slog.Error("save to database", err)
		return "", err
	}

	return CreateToken(user, os.Getenv("SECRET_KEY"))
}
