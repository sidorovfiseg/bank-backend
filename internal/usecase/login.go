package usecase

import (
	"bank-backend/internal/domain"
	"context"
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
		return user, err
}