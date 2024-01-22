package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"log/slog"
	"net/http"
)

type POSTCreateUserHandler struct {
	createUseCase *usecase.CreateUserUseCase
}

func NewPOSTUserHandler(createUseCase *usecase.CreateUserUseCase) *POSTCreateUserHandler {
	return &POSTCreateUserHandler{
		createUseCase: createUseCase,
	}
}

type POSTCreateRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}

type POSTCreateResponse struct {
	Token []byte `json:"token"`
}


// create user handler
func (h *POSTCreateUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTCreateRequest
	
	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		slog.Error("Bad request", err)
		return
	}

	command := &usecase.CreateCommand {
		Login: body.Login, 
		Email: body.Email,
		Password: []byte(body.Password),
	}

	token, err := h.createUseCase.CreateUserHandler(request.Context(), command)

	if err != nil {
		slog.Error("User creation", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("User creation", token)
	writer.WriteHeader(http.StatusOK)
	
	

}
