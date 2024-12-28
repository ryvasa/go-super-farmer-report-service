package dto

import "github.com/google/uuid"

type SupplyCreateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id"`
	CityID      int64     `json:"city_id"`
	Quantity    float64   `json:"quantity" validate:"required,gte=0"`
}

type SupplyUpdateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id"`
	CityID      int64     `json:"city_id"`
	Quantity    float64   `json:"quantity" validate:"required,gte=0"`
}
