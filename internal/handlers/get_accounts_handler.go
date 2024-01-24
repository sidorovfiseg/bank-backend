package handlers

import (
	"bank-backend/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

//TODO make handler

type GETAccountsHandler struct {
	getAccountsUseCase *usecase.GetAccountsUseCase
}

func NewGETAccountsHandler(getAccountsUseCase *usecase.GetAccountsUseCase) *GETAccountsHandler {
	return &GETAccountsHandler{
		getAccountsUseCase: getAccountsUseCase,
	}
}

type GETAccountsResponse struct {
	id uuid.UUID
	name string 
	balance float64
	userId uuid.UUID
}

type GETAccountsParams struct {
	name string `json:"name, omitempty"`
	itemsPerPage string `json:"itemsPerPage, omitempty"`
	page string  `json:"page, omitempty"`
}

func (param *GETAccountsParams) validateName() string {
	if len(param.name) > 0 {
		return param.name
	} else {
		return ""
	}
}

func (param *GETAccountsParams) validateItemsPerPage() int {
	itemsPerPage, err := strconv.Atoi(param.itemsPerPage)

	if err != nil && itemsPerPage <= 0 {
			return 1
	} else {
		return itemsPerPage
	}
}

func (param *GETAccountsParams) validatePage() int {
	page, err := strconv.Atoi(param.page)

	if err != nil && page <= 0 {
		return 1
	} else {
		return page
	}
}

func (response *GETAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
		Balance float64 `json:"balance"`
		UserId uuid.UUID `json:"user_id"`
	}{
		ID: response.id,
		Name: response.name,
		Balance: response.balance,
		UserId: response.userId,
	})
}


func (h *GETAccountsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	
	var responses []GETAccountsResponse

	err := request.ParseForm()

	if err != nil {
		return 
	}

	params := GETAccountsParams{
		name: request.Form.Get("name"),
		itemsPerPage: request.Form.Get("itemsPerPage"),
		page: request.Form.Get("page"),
	}

	name := params.validateName()
	itemsPerPage := params.validateItemsPerPage()
	page := params.validatePage()

	accounts, err := h.getAccountsUseCase.GetAccountsHandler(request.Context(), &usecase.GetAccountsCommand{
		Name: name,
		ItemsPerPage: itemsPerPage,
		Page: page,
	})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, account := range accounts {
		response := GETAccountsResponse{
			id: account.GetId(),
			name: account.GetName(),
			balance: account.GetBalance(),
			userId: account.GetUserId(),
		}
		responses = append(responses, response)
	}

	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return 
	}

}