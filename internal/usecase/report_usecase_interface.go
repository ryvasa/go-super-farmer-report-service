package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	pb "github.com/ryvasa/go-super-farmer-report-service/proto/generated"
)

type ReportUsecase interface {
	GetReportPrice(ctx context.Context, commodityID uuid.UUID, cityID int64, startDate, endDate time.Time) (*pb.ReportResponse, error)
	GetReportHarvest(ctx context.Context, landCommodityID uuid.UUID, startDate, endDate time.Time) (*pb.ReportResponse, error)
}
