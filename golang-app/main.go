package main

import (
	"fmt"
	"hl-hw4/api/gapi"
	"hl-hw4/api/nbu"
	"hl-hw4/service"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		currencyService := service.NewCurrencyService(
			nbu.NewClient(client),
			gapi.NewClient(client, os.Getenv("GA_MEASUREMENT_ID"), os.Getenv("GA_API_SECRET")),
		)

		currencyCode := "USD"
		exchangeRate, err := currencyService.GetExchangeRate(currencyCode)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = currencyService.SendExchangeRateToAnalytics(exchangeRate.Rate)
		if err != nil {
			fmt.Println("Error sending exchange rate to Google Analytics:", err)
			return
		}

		fmt.Printf("Currency: %s (Code: %s)\nExchange Rate: %f\nExchange Date: %s\n", exchangeRate.Txt, exchangeRate.CurrencyCode, exchangeRate.Rate, exchangeRate.ExchangeDate)

		time.Sleep(time.Minute)
	}
}
