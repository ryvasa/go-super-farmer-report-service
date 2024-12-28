package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Demand struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:varchar(36)"`
	CommodityID uuid.UUID      `gorm:"not null;type:varchar(36)"`
	Commodity   *Commodity     `gorm:"foreignKey:CommodityID" json:"commodity,omitempty"`
	CityID      int64          `gorm:"not null;type:int64"`
	City        *City          `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Quantity    float64        `gorm:"not null;type:float"`
	Unit        string         `gorm:"not null;type:varchar(255); default:kg"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
