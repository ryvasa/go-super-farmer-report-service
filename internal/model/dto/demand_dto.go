package dto

import "github.com/google/uuid"

type DemandCreateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id" validate:"required"`
	CityID      int64     `json:"city_id" validate:"required"`
	Quantity    float64   `json:"quantity" validate:"required,gte=0"`
}

type DemandUpdateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id"`
	CityID      int64     `json:"city_id"`
	Quantity    float64   `json:"quantity" validate:"gte=0"`
}
