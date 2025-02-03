//go:build wireinject
// +build wireinject

package wire_excel

import (
	"github.com/google/wire"
	"github.com/ryvasa/go-super-farmer-report-service/cmd/app"
	"github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/handler"
	"github.com/ryvasa/go-super-farmer-report-service/internal/repository"
	"github.com/ryvasa/go-super-farmer-report-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/database"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/minio"
	"github.com/ryvasa/go-super-farmer-report-service/utils"
)

var allSet = wire.NewSet(
	// Infrastructure
	env.LoadEnv,
	database.NewPostgres,
	minio.NewMinioClient,

	// Repository
	repository.NewReportRepositoryImpl,

	// Service
	usecase.NewExcelImpl,
	usecase.NewReportUsecase,

	// handler
	handler.NewReportHandler,

	// App
	app.NewApp,

	utils.NewGlobFunc,
)

func InitializeReportApp() (*app.ReportApp, error) {
	wire.Build(allSet)
	return nil, nil
}
