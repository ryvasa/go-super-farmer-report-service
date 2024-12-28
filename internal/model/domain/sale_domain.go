package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sale struct {
	ID          uuid.UUID      `gorm:"primary_key"`
	CityID      int64          `gorm:"not null"`
	City        *City          `gorm:"foreignKey:CityID" json:"city,omitempty"`
	CommodityID uuid.UUID      `gorm:"not null"`
	Commodity   *Commodity     `gorm:"foreignKey:CommodityID;references:ID"`
	Quantity    float64        `gorm:"not null"`
	Unit        string         `gorm:"not null;default:kg"`
	Price       float64        `gorm:"not null"`
	SaleDate    time.Time      `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
