//go:build wireinject
// +build wireinject

package wire_excel

import (
	"github.com/google/wire"
	"github.com/ryvasa/go-super-farmer-report-service/cmd/app"
	report_handler "github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/http/handler"
	report_route "github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/http/routes"
	"github.com/ryvasa/go-super-farmer-report-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-report-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/database"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/messages"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/minio"
	"github.com/ryvasa/go-super-farmer-report-service/utils"
)

var allSet = wire.NewSet(
	// Infrastructure
	env.LoadEnv,
	database.NewPostgres,
	messages.NewRabbitMQ,
	minio.NewMinioClient,

	// Repository
	repository.NewReportRepositoryImpl,

	// Service
	usecase.NewExcelImpl,
	usecase.NewReportUsecase,

	// Handler
	report_handler.NewReportHandler,

	// App
	app.NewApp,

	report_route.NewRoutes,
	report_handler.NewHandlers,

	utils.NewGlobFunc,
)

func InitializeReportApp() (*app.ReportApp, error) {
	wire.Build(allSet)
	return nil, nil
}
