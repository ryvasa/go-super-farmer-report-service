package dto

import (
	"time"

	"github.com/google/uuid"
)

type HarvestCreateDTO struct {
	LandCommodityID uuid.UUID `json:"land_commodity_id" validate:"required"`
	CityID          int64     `json:"city_id" validate:"required"`
	HarvestDate     string    `json:"harvest_date" validate:"required"`
	Quantity        float64   `json:"quantity" validate:"required"`
	Unit            string    `json:"unit" validate:"omitempty"`
}

type HarvestUpdateDTO struct {
	HarvestDate string  `json:"harvest_date" validate:"omitempty"`
	Quantity    float64 `json:"quantity" validate:"omitempty,gte=0"`
	Unit        string  `json:"unit" validate:"omitempty"`
}

type HarvestParamsDTO struct {
	LandCommodityID uuid.UUID `json:"land_commodity_id" validate:"required"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	EndDate         time.Time `json:"end_date" validate:"required"`
}
