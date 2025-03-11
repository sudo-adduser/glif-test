package database

import (
	"time"
)

type Transaction struct {
	Id        string    `gorm:"primary_key" json:"id"`
	Hash      string    `gorm:"uniqueIndex;not null" json:"hash"`
	Sender    string    `gorm:"not null" json:"sender"`
	Receiver  string    `gorm:"not null" json:"receiver"`
	Amount    string    `gorm:"not null" json:"amount"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	Status    string    `gorm:"default:'pending'" json:"status"`
}

var Models = []interface{}{&Transaction{}}
