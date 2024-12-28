package dto

import (
	"time"

	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
)

type ForecastsResponseDTO struct {
	CurrentPrice float64           `json:"price"`
	HarvestDate  time.Time         `json:"harvestDate"`
	HarvestPrice float64           `json:"harvestPrice"`
	Commodity    *domain.Commodity `json:"commodity"`
	City         *domain.City      `json:"city"`
}
