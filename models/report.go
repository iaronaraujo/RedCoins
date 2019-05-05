package models

import (
	"time"

	"github.com/iaronaraujo/RedCoins/utils"

	"github.com/iaronaraujo/RedCoins/lib"
)

//TransactionType represents the type of a transaction
type TransactionType string

const (
	//BuyBitCoins represents the type related to BitCoins buy order
	BuyBitCoins TransactionType = "BUY_BC"

	//SellBitCoins represents the type related to BitCoins sell order
	SellBitCoins TransactionType = "SELL_BC"
)

//Report represents a transaction report in redcoins
type Report struct {
	ID              int64              `db:"id" json:"id"`
	Transaction     TransactionType    `db:"transaction" json:"transaction"`
	BitCoins        float32            `db:"bitcoins" json:"bit_coins"`
	Value           float32            `db:"value" json:"value"`
	Currency        utils.CurrencyType `db:"currency" json:"currency"`
	TransactionDate time.Time          `db:"transaction_date" json:"transaction_date"`
	UserID          int64              `db:"user_id" json:"user_id"`
}

//ReportModel receives the Database table data related to reports
var ReportModel = lib.Sess.Collection("reports")
