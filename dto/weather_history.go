package dto

type WeatherHistoryDTO struct {
	Timestamp   int     `json:"timestamp"`
	Temperature float64 `json:"temperature"`
}
