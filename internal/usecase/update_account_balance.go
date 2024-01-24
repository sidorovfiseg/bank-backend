package usecase

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/pkg/persistence"
	"context"

	"github.com/google/uuid"
)

type UpdateBalanceCommand struct {
	Id uuid.UUID
	Amount int
}

type UpdateBalanceUseCase struct {
	accountRepository domain.AccountRepository
	transactionManager persistence.TransactionManager
}

func NewUpdateBalanceUseCase(accountRepository domain.AccountRepository, transactionManager persistence.TransactionManager) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		accountRepository: accountRepository,
		transactionManager: transactionManager,
	}
}

func (useCase *UpdateBalanceUseCase) UpdateBalanceHandler(ctx context.Context, command *UpdateBalanceCommand) error {
	
	err := useCase.transactionManager.Do(ctx, func(ctx context.Context) error {
		
		err := useCase.accountRepository.UpdateBalance(ctx, command.Id, command.Amount)
		
		return err
	})
	return err
}