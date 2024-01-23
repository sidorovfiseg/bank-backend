package handlers

import (
	"bank-backend/internal/middlewares"
	"bank-backend/internal/usecase"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type POSTCreateAccountHandler struct {
	createAccountUseCase *usecase.CreateAccountUseCase
}

func NewPOSTCreateAccountHandler(createAccountUseCase *usecase.CreateAccountUseCase) *POSTCreateAccountHandler {
	return &POSTCreateAccountHandler{
		createAccountUseCase: createAccountUseCase,
	}
}

type POSTCreateAccountRequest struct {
	Name string `json:"name"`
}

type POSTCreateAccountResponse struct {
	id uuid.UUID `json:"id"`
	name string		`json:"name"`
	balance float64 `json:"balance"`
	userId uuid.UUID `json:"user_id"`
}

func (h *POSTCreateAccountHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	
	var body POSTCreateAccountRequest
	userId := middlewares.GetIdFromContext(request.Context())

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		slog.Error("create account bad request", err)
		return 
	}

	command := &usecase.CreateAccountCommand{
		Name: body.Name,
		UserId: userId,
	}

	account, err := h.createAccountUseCase.CreateAccountHandler(request.Context(), command)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &POSTCreateAccountResponse{
		id: account.GetId(),
		name: account.GetName(),
		balance: account.GetBalance(),
		userId: account.GetUserId(),
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return 
	}

}