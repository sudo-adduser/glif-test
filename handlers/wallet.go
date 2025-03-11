package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"glif-test/common"
	"net/http"
)

func (h Handler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]

	balance, err := h.client.GetWalletBalance(address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.WriteLogger(w.Write([]byte("[handler:wallet] Invalid or malformed wallet address at GET /v1/wallet/{address}")))
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.WriteLogger(w.Write([]byte("[handler:wallet] Error formatting response at GET /v1/wallet/{address}")))
	}
}
