package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	db *gorm.DB
}

func InitDb(databaseURL string) *Postgres {

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(Models...)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return &Postgres{db}
}

func (p *Postgres) InsertTransaction(transaction Transaction) error {
	return p.db.Create(&transaction).Error
}

func (p *Postgres) UpdateTransaction(transaction Transaction) error {
	return p.db.Save(&transaction).Error
}

func (p *Postgres) GetTransactionsByAddress(address string) ([]Transaction, error) {
	var entries []Transaction
	err := p.db.Where("sender = ? OR receiver = ?", address, address).Find(&entries).Error
	return entries, err
}

func (p *Postgres) GetTransactionByHash(hash string) (Transaction, error) {
	entry := Transaction{}
	err := p.db.Where("hash = ?", hash).First(&entry).Error
	return entry, err
}
