package postgres

import (
	"bank-backend/internal/domain"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		dbpool: dbpool,
	}
}

func (userRepository *UserRepository) Save(ctx context.Context, user *domain.User) error {
	
	slog.Info("insert user in db user")

	_, err := userRepository.dbpool.Exec(ctx, `INSERT INTO users (user_id, login, email, password) 
																			VALUES($1, $2, $3, $4);`, 
																			user.GetId(), 
																			user.GetLogin(), 
																			user.GetEmail(), 
																			user.GetPasswordHash())
	
	if err != nil {
		slog.Error("insert user", err)
	}

	return err
}

// func (userRepository *UserRepository) FindByName(ctx context.Context, name string) 