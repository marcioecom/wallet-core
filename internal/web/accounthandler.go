package web

import (
	"encoding/json"
	"net/http"

	"github.com/marcioecom/wallet-core/internal/usecase/createaccount"
)

type AccountHandler struct {
	createAccountUseCase createaccount.CreateAccountUseCase
}

func NewAccountHandler(createAccountUseCase createaccount.CreateAccountUseCase) *AccountHandler {
	return &AccountHandler{
		createAccountUseCase: createAccountUseCase,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto createaccount.CreateAccountInputDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.createAccountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(output); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
