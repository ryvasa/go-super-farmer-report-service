package usecase

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
)

type Message struct {
	Service string      `json:"Service"`
	Data    interface{} `json:"Data"`
}

type ReportUsecaseImpl struct {
	reportRepo   repository.ReportRepository
	excelService ExcelInterface
}
type PriceMessage struct {
	CommodityID uuid.UUID `json:"CommodityID"`
	CityID      int64     `json:"CityID"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
}

type HarvestMessage struct {
	LandCommodityID uuid.UUID `json:"LandCommodityID"`
	StartDate       time.Time `json:"StartDate"`
	EndDate         time.Time `json:"EndDate"`
}

func NewReportUsecase(repo repository.ReportRepository, excelSvc ExcelInterface) ReportUsecase {
	return &ReportUsecaseImpl{
		reportRepo:   repo,
		excelService: excelSvc,
	}
}

func (u *ReportUsecaseImpl) HandlePriceHistoryMessage(msgBody []byte) error {
	var msg PriceMessage
	if err := json.Unmarshal(msgBody, &msg); err != nil {
		return err
	}

	// Ambil data dari repository
	results, err := u.reportRepo.GetPriceHistoryReport(msg.StartDate, msg.EndDate, msg.CommodityID, msg.CityID)
	if err != nil {
		return err
	}

	// Generate excel menggunakan usecase
	if err := u.excelService.CreatePriceHistoryReport(results, results[0].Commodity.Name, results[0].City.Name, results[0].Commodity.ID, results[0].CityID, msg.StartDate, msg.EndDate); err != nil {
		return err
	}

	return nil
}

func (u *ReportUsecaseImpl) HandleHarvestMessage(msgBody []byte) error {
	var msg HarvestMessage
	logrus.Log.Info("from rabbitmq usecase")
	if err := json.Unmarshal(msgBody, &msg); err != nil {
		logrus.Log.Error(err)
		return err
	}

	// Ambil data dari repository
	results, err := u.reportRepo.GetHarvestReport(msg.StartDate, msg.EndDate, msg.LandCommodityID)
	if err != nil {
		logrus.Log.Error(err)
		return err
	}

	if len(results) == 0 {
		logrus.Log.Error("no data found")
		return err
	}

	// Generate excel menggunakan usecase
	if err := u.excelService.CreateHarvestReport(results, results[0].LandCommodity.Commodity.Name, results[0].City.Name, results[0].LandCommodity.Land.User.Name, results[0].LandCommodity.ID, msg.StartDate, msg.EndDate); err != nil {
		logrus.Log.Error(err)
		return err
	}

	return nil
}
