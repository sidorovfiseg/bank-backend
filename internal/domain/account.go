package domain

import (
	"context"

	"github.com/google/uuid"
)

type Account struct {
	id uuid.UUID
	name string
	balance float64
	userId uuid.UUID
}

func NewAccount(id uuid.UUID, name string, balance float64, userId uuid.UUID) *Account {
	return &Account{
		id: id,
		name: name,
		balance: balance,
		userId: userId,
	}
}

func (account *Account) SetId(id uuid.UUID) {
	account.id = id
}

func (account *Account) SetName(name string) {
	account.name = name
}

func (account *Account) SetBalance(balance float64) {
	account.balance = balance
}

func (account *Account) SetUserId(userId uuid.UUID) {
	account.userId = userId
}

func (account *Account) GetId() uuid.UUID {
	return account.id
}

func (account *Account) GetName() string {
	return account.name
}

func (account *Account) GetBalance() float64 {
	return account.balance
}

func (account *Account) GetUserId() uuid.UUID {
	return account.userId
}

type AccountRepository interface{
	Save(ctx context.Context, account *Account) error
	FindByName(ctx context.Context, name string, userId uuid.UUID) (*Account, error)
	FindAccountsWithFilter(ctx context.Context, name string, userId uuid.UUID, itemsPerPage int, page int) ([]Account, error)
}