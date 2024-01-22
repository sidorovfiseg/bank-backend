package domain

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	id uuid.UUID
	login string
	passwordHash []byte
}

func NewUser(id uuid.UUID, login string, passwordHash []byte) *User {
	return &User{
		id: id,
		login: login,
		passwordHash: passwordHash,
	}
}

func (user *User) SetId(id uuid.UUID) {
	user.id = id
}

func (user *User) SetLogin(login string) {
	user.login = login
}

func (user *User) SetPasswordHash(passwordHash []byte) {
	user.passwordHash = passwordHash
}

func (user *User) GetId() uuid.UUID {
	return user.id
}

func (user *User) GetLogin() string {
	return user.login
}

func (user *User) GetPasswordHash() []byte {
	return user.passwordHash
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByName(ctx context.Context, login string) (*User, error)
}