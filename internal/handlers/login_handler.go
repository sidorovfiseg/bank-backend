package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"log/slog"
	"net/http"
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

// Обработка запроса на аутентифыикацию, возвращение токена 
func (h *POSTLoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTLoginRequest

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		slog.Error("login bad request", err)
		return 
	}

	token, err := h.useCase.LoginHandler(request.Context(), &usecase.LoginCommand{
		Login: body.Login,
		Password: []byte (body.Password),
	})

	if err != nil {
		slog.Error("login incorrect credentials")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Authorization", token)
	writer.WriteHeader(http.StatusOK)
	

}