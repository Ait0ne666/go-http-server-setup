package api

import (
	"belster/internal/config"
	infrastruct "belster/pkg/infrastruct"
	"belster/pkg/logger"
	"encoding/json"
	"net/http"
)

type Handlers struct {
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{}
}

func (h *Handlers) Ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func apiErrorEncode(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if customError, ok := err.(*infrastruct.CustomError); ok {
		w.WriteHeader(customError.Code)
	}

	result := struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		logger.LogError(err)
	}
}

func apiResponseEncoder(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.LogError(err)
	}
}
