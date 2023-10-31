package nbu

import (
	"context"
	"encoding/json"
	"fmt"
	"hl-hw4/entity"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://bank.gov.ua"

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{client: client}
}

func (c *Client) GetCurrencyByCode(currencyCode string) (entity.ExchangeRate, error) {
	url := fmt.Sprintf("%s/NBUStatService/v1/statdirectory/exchange?json", baseURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return entity.ExchangeRate{}, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return entity.ExchangeRate{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return entity.ExchangeRate{}, fmt.Errorf("API request returned status code %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.ExchangeRate{}, err
	}

	var exchangeRates []entity.ExchangeRate
	if err := json.Unmarshal(responseBody, &exchangeRates); err != nil {
		return entity.ExchangeRate{}, err
	}

	for _, rate := range exchangeRates {
		if rate.CurrencyCode == currencyCode {
			return rate, nil
		}
	}

	return entity.ExchangeRate{}, fmt.Errorf("currency with cc=%s not found in the API response", currencyCode)
}
