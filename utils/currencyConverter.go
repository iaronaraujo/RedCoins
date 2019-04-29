package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

//CurrencyType represents a currency used in a country
type CurrencyType string

const (
	//BRL represents brazil's currency
	BRL CurrencyType = "BRL"
	//USD represents USA's currency
	USD CurrencyType = "USD"
)

var currencyDataInst currencyData

func init() {
	currencyDataInst.bitcoinQuoteMap = make(map[CurrencyType]float32)
	updateCurrencyData()
}

type currencyData struct {
	lastUpdated     time.Time
	bitcoinQuoteMap map[CurrencyType]float32
}

type bitcoinData struct {
	Data struct {
		Quotes struct {
			BRL struct {
				Price float64 `json:"price"`
			} `json:"BRL"`
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quotes"`
	} `json:"data"`
}

// falta add tratamento de exceção aqui
func updateCurrencyData() {
	response, _ := http.Get("https://api.coinmarketcap.com/v2/ticker/1/?convert=BRL")
	var bcdata bitcoinData
	err := json.NewDecoder(response.Body).Decode(&bcdata)
	if err != nil {
		println(err.Error())
	}
	currencyDataInst.lastUpdated = time.Now()
	currencyDataInst.bitcoinQuoteMap[BRL] = float32(bcdata.Data.Quotes.BRL.Price)
	currencyDataInst.bitcoinQuoteMap[USD] = float32(bcdata.Data.Quotes.USD.Price)
}

/*
getBitCoinQuote gets the quote for a certain currency.
Returns how much 1 bitcoin can buy of a certain currency.
The quote is returned with an 1 hour margin of error */
func getBitCoinQuote(currTyp CurrencyType) float32 {
	now := time.Now()
	limit := currencyDataInst.lastUpdated.Add(time.Hour)
	if now.After(limit) {
		updateCurrencyData()
	}
	return currencyDataInst.bitcoinQuoteMap[currTyp]
}

//ConvertBitcoinsToCurrency converts an amount of bitcoins to a certain currency defined by currTyp
func ConvertBitcoinsToCurrency(bitcoins float32, currTyp CurrencyType) float32 {
	return bitcoins * getBitCoinQuote(currTyp)
}
