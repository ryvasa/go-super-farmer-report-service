package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/dto"
)

type ExcelInterface interface {
	CreatePriceHistoryReport(results []domain.PriceHistory, commodityName, regionName string, commodityID uuid.UUID, cityID int64, startDate, endDate time.Time) (string, error)
	CreateHarvestReport(results []domain.Harvest, commodityName, regionName, farmerName string, commodityID uuid.UUID, startDate, endDate time.Time) error
	GetPriceExcelFile(ctx context.Context, params *dto.PriceParamsDTO) (*string, error)
	GetHarvestExcelFile(ctx context.Context, params *dto.HarvestParamsDTO) (*string, error)
}
