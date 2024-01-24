package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type POSTUpdateBalanceHandler struct {
	useCase *usecase.UpdateBalanceUseCase
}

func NewPOSTUpdateBalanceHandler(useCase *usecase.UpdateBalanceUseCase) *POSTUpdateBalanceHandler {
	return &POSTUpdateBalanceHandler{
		useCase: useCase,
	}
}

type POSTUpdateBalanceRequest struct {
	Id uuid.UUID `json:"account_id"`
	Amount int `json:"amount"`
}

type POSTUpdateBalanceResponse struct {

}

func (h *POSTUpdateBalanceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTUpdateBalanceRequest

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	command := usecase.UpdateBalanceCommand{
		Id: body.Id,
		Amount: body.Amount,
	}

	err = h.useCase.UpdateBalanceHandler(request.Context(), &command)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)

}