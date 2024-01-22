package domain

import "github.com/google/uuid"

type Account struct {
	id uuid.UUID
	name string
	balance float64
	userId uuid.UUID
}

func NewAccount(name string, balance float64, userId uuid.UUID) *Account {
	return &Account{
		id: uuid.New(),
		name: name,
		balance: balance,
		userId: userId,
	}
}

func (account *Account) SetId(id uuid.UUID) {
	account.id = id
}

func (account *Account) setName(name string) {
	account.name = name
}

func (account *Account) setBalance(balance float64) {
	account.balance = balance
}

func (account *Account) setUserId(userId uuid.UUID) {
	account.userId = userId
}

// get account id
func (account *Account) getId() uuid.UUID {
	return account.id
}

func (account *Account) getName() string {
	return account.name
}

func (account *Account) getBalance() float64 {
	return account.balance
}

func (account *Account) getUserId() uuid.UUID {
	return account.userId
}