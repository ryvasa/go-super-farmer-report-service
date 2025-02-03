package app

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	report_handler "github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/http/handler"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/messages"
	"gorm.io/gorm"
)

type ReportApp struct {
	Router   *gin.Engine
	Env      *env.Env
	DB       *gorm.DB
	RabbitMQ messages.RabbitMQ
	Handler  *report_handler.Handlers
	MinIO    *minio.Client
}

func NewApp(
	router *gin.Engine,
	env *env.Env,
	db *gorm.DB,
	rabbitMQ messages.RabbitMQ,
	handler *report_handler.Handlers,
	minioClient *minio.Client,
) *ReportApp {
	return &ReportApp{
		Router:   router,
		Env:      env,
		DB:       db,
		RabbitMQ: rabbitMQ,
		Handler:  handler,
		MinIO:    minioClient,
	}
}
