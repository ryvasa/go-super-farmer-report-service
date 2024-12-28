package report_handler

import "github.com/gin-gonic/gin"

type ReportHandler interface {
	ConsumerHandler() error
	GetPriceHistoryReportFile(c *gin.Context)
	GetHarvestReportFile(c *gin.Context)
}
