package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WeatherHistory struct {
	gorm.Model

	Temperature decimal.Decimal `gorm:"type:numeric"`
	Timestamp   int64           `gorm:"type:bigint"`
}
