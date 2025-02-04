package app

import (
	"fmt"
	"log"
	"net"

	"github.com/minio/minio-go/v7"
	"github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/handler"
	"github.com/ryvasa/go-super-farmer-report-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	pb "github.com/ryvasa/go-super-farmer-report-service/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type ReportApp struct {
	Env        *env.Env
	DB         *gorm.DB
	MinIO      *minio.Client
	GRPCServer *grpc.Server
}

func NewApp(
	env *env.Env,
	db *gorm.DB,
	minioClient *minio.Client,
	reportUsecase usecase.ReportUsecase,
) *ReportApp {
	// Inisialisasi server gRPC
	grpcServer := grpc.NewServer()
	reportHandler := handler.NewReportHandler(reportUsecase)

	// Register gRPC service
	pb.RegisterReportServiceServer(grpcServer, reportHandler)

	// Register reflection (opsional)
	reflection.Register(grpcServer)

	return &ReportApp{
		Env:        env,
		DB:         db,
		MinIO:      minioClient,
		GRPCServer: grpcServer,
	}
}

// StartGRPCServer menjalankan server gRPC
func (app *ReportApp) StartGRPCServer() {
	port := "50051" // Ambil port dari env atau default 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("gRPC server running on port %s", port)
	if err := app.GRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
