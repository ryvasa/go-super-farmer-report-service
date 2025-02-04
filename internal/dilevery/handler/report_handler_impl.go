package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	pb "github.com/ryvasa/go-super-farmer-report-service/proto/generated"
)

type ReportHandlerImpl struct {
	pb.UnimplementedReportServiceServer
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(reportUsecase usecase.ReportUsecase) *ReportHandlerImpl {
	return &ReportHandlerImpl{
		reportUsecase: reportUsecase,
	}
}

// Implementasi method GeneratePriceHistoryReport
func (h *ReportHandlerImpl) GetReportPrice(ctx context.Context, req *pb.PriceParams) (*pb.ReportResponse, error) {
	// Konversi UUID dari string ke type uuid.UUID
	logrus.Log.Info(req)
	commodityID, err := uuid.Parse(req.CommodityId)
	if err != nil {
		return nil, fmt.Errorf("invalid commodity ID: %v", err)
	}

	// Konversi waktu dari string ke time.Time
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	// Panggil use case untuk generate laporan
	fileURL, err := h.reportUsecase.GetReportPrice(ctx, commodityID, req.CityId, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate report: %v", err)
	}

	// Return response dengan file URL
	return fileURL, nil
}

func (h *ReportHandlerImpl) GetReportHarvest(ctx context.Context, req *pb.HarvestParams) (*pb.ReportResponse, error) {
	// Konversi UUID dari string ke type uuid.UUID
	logrus.Log.Info("test")
	landCommodityID, err := uuid.Parse(req.LandCommodityId)
	if err != nil {
		return nil, fmt.Errorf("invalid land commodity ID: %v", err)
	}

	// Konversi waktu dari string ke time.Time
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	// Panggil use case untuk generate laporan
	fileURL, err := h.reportUsecase.GetReportHarvest(ctx, landCommodityID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate report: %v", err)
	}

	// Return response dengan file URL
	return fileURL, nil
}
