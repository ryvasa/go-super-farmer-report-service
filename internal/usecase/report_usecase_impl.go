package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	pb "github.com/ryvasa/go-super-farmer-report-service/proto/generated"
)

type ReportServer struct {
	pb.UnimplementedReportServiceServer
}

type Message struct {
	Service string      `json:"Service"`
	Data    interface{} `json:"Data"`
}

type ReportUsecaseImpl struct {
	reportRepo   repository.ReportRepository
	excelService ExcelInterface
}

type PriceParams struct {
	CommodityID uuid.UUID `json:"CommodityID"`
	CityID      int64     `json:"CityID"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
}

type HarvestParams struct {
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

func (u *ReportUsecaseImpl) GetReportPrice(ctx context.Context, commodityID uuid.UUID, cityID int64, startDate, endDate time.Time) (*pb.ReportResponse, error) {
	// Ambil data dari repository
	results, err := u.reportRepo.GetPriceHistoryReport(startDate, endDate, commodityID, cityID)
	if err != nil {
		return nil, err
	}

	// Generate excel menggunakan usecase
	url, err := u.excelService.CreatePriceHistoryReport(results, results[0].Commodity.Name, results[0].City.Name, results[0].Commodity.ID, results[0].CityID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &pb.ReportResponse{
		ReportUrl: url,
	}, nil
}

func (u *ReportUsecaseImpl) GetReportHarvest(ctx context.Context, landCommodityID uuid.UUID, startDate, endDate time.Time) (*pb.ReportResponse, error) {
	// Ambil data dari repository
	results, err := u.reportRepo.GetHarvestReport(startDate, endDate, landCommodityID)
	if err != nil {
		logrus.Log.Error(err)
		return nil, err
	}

	if len(results) == 0 {
		logrus.Log.Error("no data found")
		return nil, err
	}
	url := "test"
	return &pb.ReportResponse{
		ReportUrl: url,
	}, nil
}
