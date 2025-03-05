package dataprovider

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type EthDataProvider struct {
	baseAddress  string
	proxyAddress string
	proxyPort    string
	apiKey       string
}

func New(baseAddress string, proxyAddress string, proxyPort string, apiKey string) *EthDataProvider {
	return &EthDataProvider{
		baseAddress:  baseAddress,
		proxyAddress: proxyAddress,
		proxyPort:    proxyPort,
		apiKey:       apiKey,
	}
}

type EtherScanTransactionList struct {
	Transactions []EtherScanTransaction `json:"result"`
}

type EtherScanTransaction struct {
	TimeStamp string `json:"timeStamp"`
	From      string `json:"from"`
	Symbol string 
	To        string `json:"to"`
	Value     string `json:"value"`
}

func (d *EthDataProvider) GetListTransactions(ctx context.Context, address string) ([]Transactions, error) {
	url, err := url.ParseRequestURI(d.baseAddress)
	if err != nil {
		// ignore
	}

	url.Query().Add("chainid", "1")
	url.Query().Add("module", "account")
	url.Query().Add("action", "txlist")
	url.Query().Add("address", address)
	url.Query().Add("startblock", "0")
	url.Query().Add("endblock", "99999999")
	url.Query().Add("apiKey", d.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		// ignore
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var txList EtherScanTransactionList
	json.Unmarshal(body, &txList)

	result := make([]Transaction, 0)
	for _, tx := range txList.Transactions {
		amount, err := strconv.ParseFloat(tx.Value, 0)
		if err != nil {
			//ignored
		}

		timestamp := strconv.ParseInt(amount)

		t := NewTransaction(tx.From,tx.To,tx.TimeStamp, "eth", amount / math.Pow(10, 18))
	}
}
