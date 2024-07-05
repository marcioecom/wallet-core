package web

import (
	"encoding/json"
	"net/http"

	"github.com/marcioecom/wallet-core/internal/usecase/createtransaction"
)

type TransactionHandler struct {
	createTransactionUseCase createtransaction.CreateTransactionUseCase
}

func NewTransactionHandler(createTransactionUseCase createtransaction.CreateTransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		createTransactionUseCase: createTransactionUseCase,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createtransaction.CreateTransactionInputDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.createTransactionUseCase.Execute(dto)
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
