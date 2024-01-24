package postgres

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/pkg/persistence"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type AccountRepository struct {
	connection persistence.Connection
}

func NewAccountRepository(connection persistence.Connection) *AccountRepository {
	return &AccountRepository{
		connection: connection,
	}
}

func (accountRepository *AccountRepository) Save(ctx context.Context, account *domain.Account) error {
	slog.Info("creating account")

	_, err := accountRepository.connection.Exec(ctx, `INSERT INTO accounts (account_id, name, balance, user_id) 
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
		balance int
	)
	slog.Info("searching account")

	err := accountRepository.connection.QueryRow(ctx, `SELECT (account_id, balance) 
	FROM accounts WHERE name=$1 AND user_id=$2;`, 
	name, userId).Scan(&id, &balance)

	if err != nil {
		slog.Error("find account", err)
		return nil, err
	}

	account := domain.NewAccount(id, name, balance, userId)

	return account, err
}

func (accountRepository *AccountRepository) FindAccountById(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	var (
		name string
		balance int
		user_id uuid.UUID
	)
	slog.Info("searching account by id")

	err := accountRepository.connection.QueryRow(ctx, `SELECT (name, balance, user_id) FROM accounts 
	WHERE account_id=$1;`, id).Scan(&name, &balance, &user_id)

	if err != nil {
		slog.Error("find account by id", err)
		return nil, err
	}

	account := domain.NewAccount(id, name, balance, user_id)

	return account, err
}

func (accountRepository *AccountRepository) FindAccountsWithFilter(ctx context.Context, name string, userId uuid.UUID, itemsPerPage int, page int) ([]domain.Account, error) {
	
	var accounts []domain.Account
	slog.Info("searching accounts")
	
	rows, err := accountRepository.connection.Query(ctx, `SELECT (account_id, name, balance) FROM accounts WHERE user_id=$1 AND name LIKE %$2% LIMIT $3 OFFSET $4;`, userId, name, itemsPerPage, page)

	if err != nil {
		slog.Error("seacrching accounts", err)
		return nil, err
	}

	for rows.Next() {
		var (
			id uuid.UUID
			balance int
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

func (accountRepository *AccountRepository) UpdateBalance(ctx context.Context, id uuid.UUID, amount int) error {
	_, err := accountRepository.connection.Exec(ctx, "UPDATE account SET balance = balance + $2 WHERE account_id=$1;", id, amount)
	if err != nil {
		return err
	}
	return err
}