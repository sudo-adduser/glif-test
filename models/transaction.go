package models

type Transaction struct {
	ID        int64  `json:"id"`
	Hash      string `json:"hash"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

type TransactionRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   string `json:"amount"`
}
