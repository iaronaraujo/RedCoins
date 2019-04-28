package models

import (
	"time"

	"github.com/iaronaraujo/RedCoins/lib"
)

//TransactionType represents the type of a transaction
type TransactionType string

const (
	//BuyBitCoins represents the type related to bit coins buy order
	BuyBitCoins TransactionType = "Compra de BC"

	//SellBitCoins represents the type related to bit coins sell order
	SellBitCoins TransactionType = "Venda de BC"
)

//Report represents a transaction report in redcoins
type Report struct {
	ID              int             `db:"id" json:"id"`
	Transaction     TransactionType `db:"transaction" json:"transaction"`
	BitCoins        float32         `db:"bitcoins" json:"bit_coins"`
	Reais           float32         `db:"reais" json:"reais"`
	TransactionDate time.Time       `db:"transaction_date" json:"transaction_date"`
	UserID          int             `db:"user_id" json:"user_id"`
}

//ReportModel receives the Database table
var ReportModel = lib.Sess.Collection("reports")
