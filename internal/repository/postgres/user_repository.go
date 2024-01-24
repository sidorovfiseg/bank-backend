package postgres

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/pkg/persistence"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type UserRepository struct {
	connection persistence.Connection
}

func NewUserRepository(connection persistence.Connection) *UserRepository {
	return &UserRepository{
		connection: connection,
	}
}

func (userRepository *UserRepository) Save(ctx context.Context, user *domain.User) error {
	
	slog.Info("user registration")

	_, err := userRepository.connection.Exec(ctx, `INSERT INTO users (user_id, login, password) 
																			VALUES($1, $2, $3);`, 
																			user.GetId(), 
																			user.GetLogin(), 
																			user.GetPasswordHash())
	
	if err != nil {
		slog.Error("insert user", err)
	}

	return err
}

func (userRepository *UserRepository) FindByName(ctx context.Context, login string) (*domain.User, error) {
	var (
		id uuid.UUID
		password []byte
	)

	err := userRepository.connection.QueryRow(ctx, "SELECT user_id, password FROM user WHERE login=$1;", login).Scan(&id, &password)

	if err != nil {
		slog.Error("find user", err)
		return nil, err
	}

	user := domain.NewUser(id, login, password)
	return user, err

} 