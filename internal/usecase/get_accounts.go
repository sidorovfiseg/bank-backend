package usecase

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/middlewares"
	"context"
)

type GetAccountsCommand struct {
	Name string 
	ItemsPerPage int
	Page int
}

type GetAccountsUseCase struct {
	accountRepository domain.AccountRepository
}

func NewGetAccountsUseCase(accountRepository domain.AccountRepository) *GetAccountsUseCase {
	return &GetAccountsUseCase{
		accountRepository: accountRepository,
	}
}

func (useCase *GetAccountsUseCase) GetAccountsHandler(ctx context.Context, command *GetAccountsCommand) ([]domain.Account, error) {
	accounts, err := useCase.accountRepository.FindAccountsWithFilter(ctx, 
		command.Name, middlewares.GetUserIdFromContext(ctx), command.ItemsPerPage, command.Page)

	return accounts, err
}