package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"log/slog"
	"net/http"
)

type POSTRegistrationHandler struct {
	registrationUseCase *usecase.RegistrationUseCase
}

func NewPOSTUserHandler(registrationUseCase *usecase.RegistrationUseCase) *POSTRegistrationHandler {
	return &POSTRegistrationHandler{
		registrationUseCase: registrationUseCase,
	}
}

type POSTRegistrationRequest struct {
	Login string `json:"login"`
	Password []byte `json:"password"`
}

type POSTRegistrationResponse struct {
	Token []byte `json:"token"`
}


// create user handler
func (h *POSTRegistrationHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTRegistrationRequest
	
	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		slog.Error("registration Bad request", err)
		return
	}

	command := &usecase.RegistrationCommand {
		Login: body.Login, 
		Password: []byte(body.Password),
	}

	token, err := h.registrationUseCase.RegistrationHandler(request.Context(), command)

	if err != nil {
		slog.Error("User creation", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("User creation", token)
	writer.WriteHeader(http.StatusCreated)

	response := &POSTRegistrationResponse{
		Token: []byte(token),
	}

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		slog.Error("response encode", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	
	

}
