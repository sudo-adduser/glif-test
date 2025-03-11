package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type EthClient struct {
	client *ethclient.Client
}

func NewClient(rpcURL string) *EthClient {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("[blockchain:client] could not connect to ethereum api %v", err)
	}
	return &EthClient{client: client}
}
