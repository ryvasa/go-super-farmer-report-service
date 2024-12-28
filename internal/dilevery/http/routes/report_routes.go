package report_route

import (
	"github.com/gin-gonic/gin"
	report_handler "github.com/ryvasa/go-super-farmer-report-service/internal/dilevery/http/handler"
)

type ReportRoutes struct {
	router  *gin.Engine
	handler report_handler.ReportHandler
}

func NewReportRoutes(router *gin.Engine, handler report_handler.ReportHandler) {
	router.GET("/prices/history/commodity/:commodity_id/city/:city_id/download/file", handler.GetPriceHistoryReportFile)
	router.GET("/harvests/land_commodity/:id/download/file", handler.GetHarvestReportFile)
}
