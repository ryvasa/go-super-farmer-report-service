package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PriceHistory struct {
	ID          uuid.UUID      `gorm:"primary_key;"`
	CommodityID uuid.UUID      `gorm:"not null"`
	Commodity   *Commodity     `gorm:"foreignKey:CommodityID;references:ID"`
	CityID      int64          `gorm:"not null"`
	City        *City          `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Price       float64        `gorm:"not null" validate:"gte=0"`
	Unit        string         `gorm:"not null;type:varchar(255); default:idr"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
