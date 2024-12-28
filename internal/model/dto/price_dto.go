package dto

import (
	"time"

	"github.com/google/uuid"
)

type PriceCreateDTO struct {
	CommodityID uuid.UUID `json:"commodity_id" validate:"required"`
	CityID      int64     `json:"city_id" validate:"required"`
	Price       float64   `json:"price" validate:"required,min=1"`
}

type PriceUpdateDTO struct {
	Price float64 `json:"price" validate:"required,min=1"`
}

type PriceResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Price       float64   `json:"price"`
	CommodityID uuid.UUID `json:"-"`
	Commodity   struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"-"`
	} `json:"commodity"`
	CityID int64 `json:"-"`
	City   struct {
		ID         int64 `json:"id"`
		ProvinceID int64 `json:"-"`
		Province   struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"province"`
	} `json:"city"`
}

type PriceParamsDTO struct {
	CommodityID uuid.UUID `json:"commodity_id" validate:"required"`
	CityID      int64     `json:"city_id" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
}
