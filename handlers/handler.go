package handlers

import (
	"glif-test/blockchain"
	"glif-test/database"
)

type Handler struct {
	client  *blockchain.EthClient
	db      *database.Postgres
	chainId int
}

func InitHandlers(client *blockchain.EthClient, db *database.Postgres, chainId int) *Handler {
	return &Handler{client: client, db: db, chainId: chainId}
}
