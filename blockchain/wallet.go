package blockchain

import (
	"context"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"glif-test/database"
	"glif-test/models"
	"log"
	"math"
	"math/big"
	"os"
	"time"
)

func (e EthClient) GetWalletBalance(address string) (string, error) {
	ctx := context.Background()
	hexAddress := common.HexToAddress(address)
	balance, err := e.client.BalanceAt(ctx, hexAddress, nil)
	if err != nil {
		return "", err
	}

	// Convert wei to ETH for readability and return as a string for JSON serialization
	test := balance.String()
	fmt.Println(test)
	etherBalance := new(big.Float).SetInt(balance)
	etherBalance.Quo(etherBalance, big.NewFloat(math.Pow10(18)))

	return etherBalance.String(), nil
}

func (e EthClient) SubmitTransaction(txReq models.TransactionRequest, db *database.Postgres, chainId int) (string, error) {
	// This could be handled many different ways depending on the product requirements this was just a quick and dirty way to get the test private key
	testPk := os.Getenv("TEST_PRIVATE_KEY")

	privateKey, err := crypto.HexToECDSA(testPk)
	if err != nil {
		return "", err
	}
	fromAddress := common.HexToAddress(txReq.Sender)

	toAddress := common.HexToAddress(txReq.Receiver)

	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	txValue := new(big.Int)
	txValue, ok := txValue.SetString(txReq.Amount, 10)

	if !ok {
		fmt.Println("Error converting string to big.Int")
		return "", errors.New("[blockchain:wallet] error converting transaction value string to big.Int")
	}

	txData := types.NewTx(&types.LegacyTx{
		To:       &toAddress,
		Value:    txValue,
		Gas:      21000,
		GasPrice: gasPrice,
		Nonce:    nonce,
		Data:     nil, // Optional, include data if you're calling a contract
	})

	// TODO: currently hardcoded to Holsky but this should be configurable

	signer := types.LatestSignerForChainID(big.NewInt(int64(chainId)))
	signedTx, err := types.SignTx(txData, signer, privateKey)
	if err != nil {
		return "", err
	}

	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	go e.trackTransactionStatus(signedTx.Hash(), db)

	return signedTx.Hash().Hex(), nil
}

func (e EthClient) trackTransactionStatus(txHash common.Hash, db *database.Postgres) {
	b := backoff.NewExponentialBackOff()

	// Set the initial interval to 1 minute
	b.InitialInterval = time.Minute

	tx, err := db.GetTransactionByHash(txHash.Hex())
	if err != nil {
		// TODO figure out how we want to handle error

	}

	operation := func() error {
		receipt, err := e.client.TransactionReceipt(context.Background(), txHash)

		if err != nil {
			if err.Error() == "not found" {
				fmt.Printf("Transaction %s not found yet, retrying...\n", txHash.Hex())
				return err
			}

			fmt.Printf("Error fetching receipt for transaction %s: %v\n", txHash.Hex(), err)
			return backoff.Permanent(err)
		}

		// Process the receipt
		if receipt != nil {
			switch receipt.Status {
			case types.ReceiptStatusSuccessful:
				fmt.Printf("Transaction %s successfully processed\n", txHash.Hex())
				tx.Status = "confirmed"
				err = db.UpdateTransaction(tx)
				if err != nil {
					log.Printf("Error updating transaction %s: %v\n", txHash.Hex(), err)
					return backoff.Permanent(err)
				}
			case types.ReceiptStatusFailed:
				fmt.Printf("Transaction %s failed\n", txHash.Hex())
				tx.Status = "failed"
				err = db.UpdateTransaction(tx)
				if err != nil {
					log.Printf("Error updating transaction %s: %v\n", txHash.Hex(), err)
					return backoff.Permanent(err)
				}
			default:
				fmt.Printf("Transaction %s has an unknown status: %d\n", txHash.Hex(), receipt.Status)
			}
			return nil
		}

		return nil
	}

	// Execute the backoff operation with retries
	if err := backoff.Retry(operation, b); err != nil {
		fmt.Printf("Transaction %s tracking failed: %v\n", txHash.Hex(), err)
	}
}
