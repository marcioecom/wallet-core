package web

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Respond(w http.ResponseWriter, code int, err error, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	res := httpResponse{
		Code:  code,
		Error: err.Error(),
		Data:  data,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
