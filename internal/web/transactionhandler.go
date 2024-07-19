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
		Respond(w, http.StatusBadRequest, err, nil)
		return
	}

	output, err := h.createTransactionUseCase.Execute(r.Context(), dto)
	if err != nil {
		Respond(w, http.StatusBadRequest, err, nil)
		return
	}

	Respond(w, http.StatusCreated, nil, output)
}
