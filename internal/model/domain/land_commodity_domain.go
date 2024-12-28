package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LandCommodity struct {
	ID          uuid.UUID      `gorm:"primary_key;"`
	LandArea    float64        `gorm:"not null"`
	Unit        string         `gorm:"not null;type:varchar(255); default:ha"`
	CommodityID uuid.UUID      `gorm:"not null"`
	Commodity   *Commodity     `gorm:"foreignKey:CommodityID;references:ID"`
	LandID      uuid.UUID      `gorm:"not null"`
	Land        *Land          `gorm:"foreignKey:LandID;references:ID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
