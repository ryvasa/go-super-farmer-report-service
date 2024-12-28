package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Harvest struct {
	ID              uuid.UUID      `gorm:"primaryKey;type:varchar(36)"`
	LandCommodityID uuid.UUID      `gorm:"not null;type:varchar(36)"`
	LandCommodity   *LandCommodity `gorm:"foreignKey:LandCommodityID"`
	CityID          int64          `gorm:"not null;type:int64"`
	City            *City          `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Quantity        float64        `gorm:"not null;type:float"`
	Unit            string         `gorm:"not null;type:varchar(255); default:kg"`
	HarvestDate     time.Time      `gorm:"not null;type:timestamp"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
