package api

import (
	"net/http"
	"time"

	owm "github.com/briandowns/openweathermap"
)

type OWMApiClient interface {
	GetWeatherHistory(location *owm.Coordinates) (*owm.HistoricalWeatherData, error)
}

type owmApiClient struct {
	client *http.Client
	apiKey string
}

func NewOWMApiClient(client *http.Client, apiKey string) (OWMApiClient, error) {
	if err := owm.ValidAPIKey(apiKey); err != nil {
		return nil, err
	}

	return &owmApiClient{
		client: client,
		apiKey: apiKey,
	}, nil
}

func (c *owmApiClient) GetWeatherHistory(location *owm.Coordinates) (*owm.HistoricalWeatherData, error) {
	history, err := owm.NewHistorical("C", c.apiKey)
	if err != nil {
		return nil, err
	}
	hp := &owm.HistoricalParameters{
		Start: time.Now().AddDate(0, -11, 0).Unix(),
		End:   time.Now().AddDate(0, 0, -2).Unix(),
		Cnt:   24,
	}
	if err := history.HistoryByCoord(location, hp); err != nil {
		return nil, err
	}

	return history, nil
}
