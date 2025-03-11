package router

import (
	"github.com/gorilla/mux"
	"glif-test/handlers"
)

func New(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/echo", handlers.GetEcho).Methods("GET")
	r.HandleFunc("/v1/wallet/{address}", h.GetWalletBalance).Methods("GET")
	r.HandleFunc("/v1/transaction", h.SubmitTransaction).Methods("POST")
	r.HandleFunc("/v1/transactions/{address}", h.GetTransactions).Methods("GET")

	return r
}
