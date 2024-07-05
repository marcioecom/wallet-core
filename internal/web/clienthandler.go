package web

import (
	"encoding/json"
	"net/http"

	"github.com/marcioecom/wallet-core/internal/usecase/createclient"
)

type ClientHandler struct {
	createClientUseCase createclient.CreateClientUseCase
}

func NewClientHandler(createClientUseCase createclient.CreateClientUseCase) *ClientHandler {
	return &ClientHandler{
		createClientUseCase: createClientUseCase,
	}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto createclient.CreateClientInputDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.createClientUseCase.Execute(dto)
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
