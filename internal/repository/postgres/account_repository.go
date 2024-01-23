package postgres

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	dbpool *pgxpool.Pool
}

func NewAccountRepository(dbpool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		dbpool: dbpool,
	}
}

func (accountRepository *AccountRepository) Save(ctx context.Context, account *domain.Account) error {
	slog.Info("creating account")

	_, err := accountRepository.dbpool.Exec(ctx, `INSERT INTO accounts (account_id, name, balance, user_id) 
	VALUES($1, $2, $3, $4);`,
	account.GetId(),
	account.GetName(),
	account.GetBalance(),
	account.GetUserId())

	if err != nil {
		slog.Error("insert account", err)
	}

	return err
}