package server

import (
	"belster/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *api.Handlers) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Methods(http.MethodGet).Path("/ping").HandlerFunc(h.Ping)

	// admin := router.PathPrefix("/admin").Subrouter()

	return router
}
