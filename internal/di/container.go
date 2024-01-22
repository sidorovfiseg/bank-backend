package di

import (
	"bank-backend/internal/domain"
	"bank-backend/internal/handlers"
	"bank-backend/internal/repository/postgres"
	"bank-backend/internal/usecase"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {

		router http.Handler
		dbcontainer DatabaseContainer

		createUser *usecase.RegistrationUseCase
		userRepository *postgres.UserRepository
		postCreateUserHandler *handlers.POSTRegistrationHandler
}

func NewContainer(ctx context.Context) *Container {
	return &Container{
		dbcontainer: *New(ctx),
	}
}

func (c *Container) getPool() *pgxpool.Pool {
	return c.dbcontainer.dbpool
}

// func (c *Container) SetUserRepository(userRepository domain.UserRepository) {
// 	c.userRepository = userRepository
// }

func (c *Container) GetUserRepository() domain.UserRepository {
	if (c.userRepository == nil) {
		c.userRepository = postgres.NewUserRepository(c.getPool())
	}

	return c.userRepository
}

func (c *Container) CreateUser() *usecase.RegistrationUseCase {
	if (c.createUser == nil) {
		c.createUser = usecase.NewRegistrationUseCase(c.GetUserRepository())
	}

	return c.createUser
}

func (c *Container) PostCreateUserHandler() *handlers.POSTRegistrationHandler {
	if (c.postCreateUserHandler == nil) {
		c.postCreateUserHandler = handlers.NewPOSTUserHandler(c.CreateUser())
	}

	return c.postCreateUserHandler
}

// func (c *Container) HTTPRouter() http.Handler {
// 	if c.router != nil {
// 		return c.router
// 	}

// 	router := mux.NewRouter()
	
// }