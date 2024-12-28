package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
)

type ReportRepository interface {
	GetPriceHistoryReport(start, end time.Time, commodityID uuid.UUID, cityID int64) ([]domain.PriceHistory, error)
	GetHarvestReport(start, end time.Time, landCommodityID uuid.UUID) ([]domain.Harvest, error)
}
