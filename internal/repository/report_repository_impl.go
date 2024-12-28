package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	"gorm.io/gorm"
)

type ReportRepositoryImpl struct {
	db *gorm.DB
}

func NewReportRepositoryImpl(db *gorm.DB) ReportRepository {
	return &ReportRepositoryImpl{db}
}

func (r *ReportRepositoryImpl) GetPriceHistoryReport(start, end time.Time, commodityID uuid.UUID, cityID int64) ([]domain.PriceHistory, error) {
	var results []domain.PriceHistory
	startOfDay := start.UTC()
	endOfDay := end.Add(24 * time.Hour).Add(-time.Second)
	// Query current price
	var currentPrice domain.Price
	err := r.db.
		Preload("Commodity", func(db *gorm.DB) *gorm.DB {
			return db.Omit("CreatedAt", "UpdatedAt", "DeletedAt", "Description")
		}).
		Preload("City").
		Preload("City.Province").
		Where("prices.commodity_id = ? AND prices.city_id = ? AND prices.created_at BETWEEN ? AND ?", commodityID, cityID, startOfDay, endOfDay).
		First(&currentPrice).Error
	if err != nil {
		return nil, err
	}

	// Konversi current price ke price history
	results = append(results, domain.PriceHistory{
		ID:          currentPrice.ID,
		CommodityID: currentPrice.CommodityID,
		CityID:      currentPrice.CityID,
		Price:       currentPrice.Price,
		Unit:        currentPrice.Unit,
		CreatedAt:   time.Now(),
		Commodity:   currentPrice.Commodity,
		City:        currentPrice.City,
	})

	// Query price histories
	var histories []domain.PriceHistory
	err = r.db.Preload("Commodity").
		Preload("City").
		Joins("JOIN commodities ON price_histories.commodity_id = commodities.id").
		Joins("JOIN cities ON price_histories.city_id = cities.id").
		Where("price_histories.commodity_id = ? AND price_histories.city_id = ? AND price_histories.deleted_at IS NULL AND price_histories.created_at BETWEEN ? AND ?",
			commodityID, cityID, startOfDay, endOfDay).
		Order("price_histories.created_at DESC").
		Find(&histories).Error
	if err != nil {
		return nil, err
	}

	results = append(results, histories...)

	return results, nil
}

func (r *ReportRepositoryImpl) GetHarvestReport(start, end time.Time, landCommodityID uuid.UUID) ([]domain.Harvest, error) {
	var results []domain.Harvest
	logrus.Log.Infof("GetHarvestReport: landCommodityID=%v, start=%v, end=%v", landCommodityID, start, end)

	startOfDay := start.UTC()
	endOfDay := end.Add(24 * time.Hour).Add(-time.Second)

	logrus.Log.Info("Try get db")
	err := r.db.
		Where("land_commodity_id = ? AND deleted_at IS NULL", landCommodityID).
		Preload("LandCommodity").
		Preload("City").
		Preload("LandCommodity.Commodity").
		Preload("LandCommodity.Land").
		Preload("LandCommodity.Land.User").
		Preload("City.Province").
		Order("harvests.created_at DESC").
		Where("harvests.created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Find(&results).Error

	if err != nil {
		logrus.Log.Error(err)
		return nil, err
	}

	logrus.Log.Infof("GetHarvestReport: Found %d harvest records", len(results)) // Tambahkan ini
	if len(results) == 0 {
		return []domain.Harvest{}, nil
	}

	return results, nil
}
