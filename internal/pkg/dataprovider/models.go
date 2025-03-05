package dataprovider

import "time"

type Transaction struct {
	From     string
	To       string
	DateTime time.Time
	Symbol   string
	Amount   float64
}

func NewTransaction(from string, to string, timestamp int64, symbol string, amount float64) Transaction {
	txTime := time.Unix(timestamp, 0).Local()
	return Transaction{
		From:     from,
		To:       to,
		Symbol:   symbol,
		Amount:   amount,
		DateTime: txTime,
	}
}
