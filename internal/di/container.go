package di

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/handlers"
	"bank-backend/internal/pkg/persistence"
	"bank-backend/internal/repository/postgres"
	"bank-backend/internal/usecase"
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	router http.Handler
	pool *pgxpool.Pool
	transactionManager *persistence.TransactionManager

	userRepository *postgres.UserRepository
	accountRepository *postgres.AccountRepository

	loginUseCase *usecase.LoginUseCase
	registrationUseCase *usecase.RegistrationUseCase

	createAccountUseCase *usecase.CreateAccountUseCase
	getAccountsUseCase *usecase.GetAccountsUseCase
	updateBalanceUseCase *usecase.UpdateBalanceUseCase

	postLoginHandler *handlers.POSTLoginHandler
	postRegistrationHandler *handlers.POSTRegistrationHandler

	postCreateAccountHandler *handlers.POSTCreateAccountHandler
	getAccountsHandler *handlers.GETAccountsHandler
	postUpdateBalanceHandler *handlers.POSTUpdateBalanceHandler

}

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		slog.Error("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(ctx, databaseURL)

	if err != nil {
		slog.Error("create connection pool", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		slog.Error("ping", err)
		return nil, err
	}

	return pool, err
} 

func NewContainer(ctx context.Context) *Container {
	pool, err := CreateConnection(ctx)
	if err != nil {
		slog.Error("error creating container", err)
	}
	return &Container{
		pool: pool,
	}
}

func (c *Container) Close() {
	c.pool.Close()
}

func (c *Container) UserRepository() domain.UserRepository {
	if c.userRepository == nil {
		c.userRepository = postgres.NewUserRepository(c.pool)
	}

	return c.userRepository
}

func (c *Container) LoginUseCase() *usecase.LoginUseCase {
	if c.loginUseCase == nil {
		c.loginUseCase = usecase.NewLoginUseCase(c.userRepository)
	}
	return c.loginUseCase
} 

func (c *Container) RegistrationUseCase() *usecase.RegistrationUseCase {
	if c.registrationUseCase == nil {
		c.registrationUseCase = usecase.NewRegistrationUseCase(c.userRepository)
	}
	return c.registrationUseCase
}



// type Container struct {

// 		router http.Handler
// 		dbcontainer DatabaseContainer

// 		createUser *usecase.RegistrationUseCase
// 		userRepository *postgres.UserRepository
// 		postCreateUserHandler *handlers.POSTRegistrationHandler
// }

// func NewContainer(ctx context.Context) *Container {
// 	return &Container{
// 		dbcontainer: *New(ctx),
// 	}
// }

// func (c *Container) getPool() *pgxpool.Pool {
// 	return c.dbcontainer.dbpool
// }

// // func (c *Container) SetUserRepository(userRepository domain.UserRepository) {
// // 	c.userRepository = userRepository
// // }

// func (c *Container) GetUserRepository() domain.UserRepository {
// 	if (c.userRepository == nil) {
// 		c.userRepository = postgres.NewUserRepository(c.getPool())
// 	}

// 	return c.userRepository
// }

// func (c *Container) CreateUser() *usecase.RegistrationUseCase {
// 	if (c.createUser == nil) {
// 		c.createUser = usecase.NewRegistrationUseCase(c.GetUserRepository())
// 	}

// 	return c.createUser
// }

// func (c *Container) PostCreateUserHandler() *handlers.POSTRegistrationHandler {
// 	if (c.postCreateUserHandler == nil) {
// 		c.postCreateUserHandler = handlers.NewPOSTUserHandler(c.CreateUser())
// 	}

// 	return c.postCreateUserHandler
// }

// func (c *Container) HTTPRouter() http.Handler {
// 	if c.router != nil {
// 		return c.router
// 	}

// 	router := mux.NewRouter()
	
// }