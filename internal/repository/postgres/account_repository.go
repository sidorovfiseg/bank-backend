package postgres

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"

	"github.com/google/uuid"
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

func (accountRepository *AccountRepository) FindByName(ctx context.Context, name string, userId uuid.UUID) (*domain.Account, error) {
	
	var (
		id uuid.UUID
		balance float64
	)
	slog.Info("searching account")

	err := accountRepository.dbpool.QueryRow(ctx, `SELECT (account_id, balance) 
	FROM accounts WHERE name=$1 AND user_id=$2`, 
	name, userId).Scan(&id, &balance)

	if err != nil {
		slog.Error("find account", err)
		return nil, err
	}

	account := domain.NewAccount(id, name, balance, userId)

	return account, err
}

func (accountRepository *AccountRepository) FindAccountsWithFilter(ctx context.Context, name string, userId uuid.UUID, itemsPerPage int, page int) ([]domain.Account, error) {
	
	var accounts []domain.Account
	slog.Info("searching accounts")
	
	rows, err := accountRepository.dbpool.Query(ctx, `SELECT (account_id, name, balance) FROM accounts WHERE user_id=$1 AND name LIKE %$2% LIMIT $3 OFFSET $4;`, userId, name, itemsPerPage, page)

	if err != nil {
		slog.Error("seacrching accounts", err)
		return nil, err
	}

	for rows.Next() {
		var (
			id uuid.UUID
			balance float64
		)
		err = rows.Scan(
			&id,
			&name,
			&balance,
		)
		if err != nil {
			slog.Error("searching account row", err)
			return nil, err
		}
		account := domain.NewAccount(id, name, balance, userId)
		accounts = append(accounts, *account)
	}

	return accounts, err
}

func (accountRepository *AccountRepository) UpdateBalance(ctx context.Context, id uuid.UUID, newBalance int) error {
	_, err := accountRepository.dbpool.Exec(ctx, "UPDATE account SET balance = $2 WHERE account_id=$1", id, newBalance)
	if err != nil {
		return err
	}
	return err
}