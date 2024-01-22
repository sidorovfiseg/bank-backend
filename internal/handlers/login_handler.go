package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5/request"
)

type POSTLoginHandler struct {
	useCase *usecase.LoginUseCase
}

func NewPOSTLoginHandler(useCase *usecase.LoginUseCase) *POSTLoginHandler {
	return &POSTLoginHandler{
		useCase: useCase,
	}
}

type POSTLoginRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type POSTLoginResponse struct {
	Token []byte `json:"token"`
}

func (h *POSTLoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTLoginRequest

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		slog.Error("login bad request", err)
		return 
	}

	

}