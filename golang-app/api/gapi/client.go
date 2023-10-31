package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://www.google-analytics.com/mp/collect"

type Client struct {
	httpClient      *http.Client
	gaMeasurementID string
	apiSecret       string
}

func NewClient(httpClient *http.Client, gaMeasurementID, apiSecret string) *Client {
	return &Client{
		httpClient:      httpClient,
		gaMeasurementID: gaMeasurementID,
		apiSecret:       apiSecret,
	}
}

func (c *Client) SendExchangeRateToAnalytics(exchangeRate float64) error {
	url := fmt.Sprintf("%s?measurement_id=%s&api_secret=%s", baseURL, c.gaMeasurementID, c.apiSecret)

	eventData := map[string]interface{}{
		"client_id": "test-146332",
		"events": []map[string]interface{}{
			{
				"name": "currency_exchange",
				"params": map[string]interface{}{
					"engagement_time_msec": "100",
					"session_id":           "123",
					"uah_usd_ratio":        exchangeRate,
				},
			},
		},
	}

	eventJSON, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(eventJSON))
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error sending data to Google Analytics. Status code: %s", resp.Status)
	}

	return nil
}
