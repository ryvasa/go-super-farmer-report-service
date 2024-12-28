package dto

import "github.com/google/uuid"

type SaleCreateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id" validate:"required"`
	CityID      int64     `json:"city_id" validate:"required"`
	Quantity    float64   `json:"quantity" validate:"required,gte=0"`
	Unit        string    `json:"unit" validate:"required"`
	Price       float64   `json:"price" validate:"required,gte=0"`
	SaleDate    string    `json:"sale_date,omitempty" validate:"omitempty"`
}

type SaleUpdateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id,omitempty" validate:"omitempty"`
	CityID      int64     `json:"city_id,omitempty" validate:"omitempty"`
	Quantity    float64   `json:"quantity,omitempty" validate:"omitempty"`
	Unit        string    `json:"unit,omitempty" validate:"omitempty"`
	Price       float64   `json:"price,omitempty" validate:"omitempty"`
	SaleDate    string    `json:"sale_date,omitempty" validate:"omitempty"`
}
