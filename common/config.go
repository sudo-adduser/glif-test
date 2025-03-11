package common

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	ServerPort int    `env:"PORT" default:"8080"`
	DbAddress  string `env:"DB_ADDRESS,required"`
	EthUrl     string `env:"ETH_URL,required"`
	ChainId    int    `env:"CHAIN_ID,required"`
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("WARNING: No .env file found, using system environment variables")
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}
