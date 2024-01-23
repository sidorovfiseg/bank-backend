package usecase

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type CreateAccountCommand struct {
	Name string
	UserId uuid.UUID
}

type CreateAccountUseCase struct {
	accountRepository domain.AccountRepository
}

func NewCreateAccountUseCase(accountRepository domain.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountRepository: accountRepository,
	}
}

func (useCase *CreateAccountUseCase) CreateAccountHandler(ctx context.Context, command *CreateAccountCommand) (*domain.Account, error) {
	account := domain.NewAccount(uuid.New(), command.Name, 0.0, command.UserId)

	err := useCase.accountRepository.Save(ctx, account)

	if err != nil {
		slog.Error("save account to db", err)
		return nil, err
	}
	
	return account, err
}