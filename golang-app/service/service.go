package service

import (
	"hl-hw4/entity"
)

type ExchangeRateClientInterface interface {
	GetCurrencyByCode(currencyCode string) (entity.ExchangeRate, error)
}

type GoogleAnalyticsClientInterface interface {
	SendExchangeRateToAnalytics(exchangeRate float64) error
}

type CurrencyService struct {
	exchangeRateClient ExchangeRateClientInterface
	gaClient           GoogleAnalyticsClientInterface
}

func NewCurrencyService(
	exchangeRateClient ExchangeRateClientInterface,
	gaClient GoogleAnalyticsClientInterface,
) *CurrencyService {
	return &CurrencyService{
		exchangeRateClient: exchangeRateClient,
		gaClient:           gaClient,
	}
}

func (s *CurrencyService) GetExchangeRate(currencyCode string) (entity.ExchangeRate, error) {
	return s.exchangeRateClient.GetCurrencyByCode(currencyCode)
}

func (s *CurrencyService) SendExchangeRateToAnalytics(exchangeRate float64) error {
	return s.gaClient.SendExchangeRateToAnalytics(exchangeRate)
}
