package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"glif-test/common"
	"glif-test/database"
	"glif-test/models"
	"net/http"
)

func (h Handler) SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var txReq models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&txReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	txHash, err := h.client.SubmitTransaction(txReq, h.db, h.chainId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to submit transatcion %s", err), http.StatusInternalServerError)
	}
	//// Store transaction in DB
	tx := database.Transaction{
		Id:       uuid.New().String(),
		Hash:     txHash,
		Sender:   txReq.Sender,
		Receiver: txReq.Receiver,
		Amount:   txReq.Amount,
	}
	err = h.db.InsertTransaction(tx)
	if err != nil {

	}
	//// Return response
	response := map[string]string{"txHash": txHash, "status": "submitted"}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (h Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]

	txs, err := h.db.GetTransactionsByAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(txs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		common.WriteLogger(w.Write([]byte("[handler:wallet] Error formatting response at GET /v1/wallet/{address}")))
	}
}
