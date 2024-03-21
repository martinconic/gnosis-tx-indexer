package indexer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Transaction struct {
	From      string `json:"from"`
	TimeStamp string `json:"timeStamp"`
}

type APITransactionResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

func GetTransactionByAddress(address, apiKey string) ([]Transaction, error) {

	requestURL := fmt.Sprintf("https://api.gnosisscan.io/api?module=account&action=txlist&address=%s&sort=desc&apikey=%s", address, apiKey)

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading response body: %s\n", err)
		return nil, err
	}

	var transactionsResponse APITransactionResponse
	err = json.Unmarshal(body, &transactionsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err)
	}

	return transactionsResponse.Result, nil
}
