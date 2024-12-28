package dto

import "github.com/google/uuid"

type LandCommodityCreateDTO struct {
	LandID      uuid.UUID `json:"land_id"`
	CommodityID uuid.UUID `json:"commodity_id"`
	LandArea    float64   `json:"land_area" validate:"required,min=1"`
}

type LandCommodityUpdateDTO struct {
	LandID      uuid.UUID `json:"land_id"`
	CommodityID uuid.UUID `json:"commodity_id"`
	LandArea    float64   `json:"land_area" validate:"required,min=1"`
}
