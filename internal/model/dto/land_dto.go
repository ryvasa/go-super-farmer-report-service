package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
)

type LandCreateDTO struct {
	LandArea    float64 `json:"land_area" validate:"required,min=1,max=10000"`
	Certificate string  `json:"certificate" validate:"required,min=1,max=255"`
}

type LandUpdateDTO struct {
	LandArea    float64 `json:"land_area,omitempty" validate:"omitempty,min=1,max=10000"`
	Certificate string  `json:"certificate,omitempty" validate:"omitempty,min=1,max=255"`
}

type LandResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	LandArea    float64   `json:"land_area"`
	Certificate string    `json:"certificate"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LandAreaParamsDTO struct {
	CommodityID uuid.UUID `json:"commodity_id" validate:"gte=0,omitempty" form:"commodity_id"`
	CityID      int64     `json:"city_id" validate:"gte=0,omitempty" form:"city_id"`
}

type LandAreaResponseDTO struct {
	LandArea       float64           `json:"land_area"`
	LandAreaActive float64           `json:"land_area_active"`
	Unit           string            `json:"unit"`
	City           *domain.City      `json:"city,omitempty"`
	Commodity      *domain.Commodity `json:"commodity,omitempty"`
}
