package main

import (
	"fmt"
	"glif-test/blockchain"
	"glif-test/handlers"
	"log"
	"net/http"
	"strconv"

	"glif-test/common"
	"glif-test/database"
	"glif-test/router"
)

func main() {
	// Load config
	cfg := common.LoadConfig()

	// Connect to database
	db := database.InitDb(cfg.DbAddress)

	// Initialize Eth Client
	c := blockchain.NewClient(cfg.EthUrl)

	// Initialize handlers
	h := handlers.InitHandlers(c, db, cfg.ChainId)

	// Setup Router
	r := router.New(h)

	// Define server address
	serverAddress := ":" + strconv.Itoa(cfg.ServerPort)

	// Start server
	fmt.Printf("Server running on port %d", cfg.ServerPort)
	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
