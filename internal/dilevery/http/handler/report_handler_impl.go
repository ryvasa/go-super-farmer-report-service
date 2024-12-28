package report_handler

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/dto"
	"github.com/ryvasa/go-super-farmer-report-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/messages"
	"github.com/ryvasa/go-super-farmer-report-service/utils"
)

type ReportHandlerImpl struct {
	reportUsecase usecase.ReportUsecase
	excelUsecase  usecase.ExcelInterface
	rabbitMQ      messages.RabbitMQ
}

func NewReportHandler(rabbitMQUsecase usecase.ReportUsecase, excelSvc usecase.ExcelInterface, rabbitMQ messages.RabbitMQ) ReportHandler {
	return &ReportHandlerImpl{rabbitMQUsecase, excelSvc, rabbitMQ}
}

func (h *ReportHandlerImpl) ConsumerHandler() error {
	// Start consumers for each queue
	err := h.startConsumer("price-history-queue", h.reportUsecase.HandlePriceHistoryMessage)
	if err != nil {
		logrus.Log.Fatal("failed to start price-history-queue consumer", err)
	}

	err = h.startConsumer("harvest-queue", h.reportUsecase.HandleHarvestMessage)
	if err != nil {
		logrus.Log.Fatal("failed to start harvest-queue consumer", err)
	}

	logrus.Log.Info("Consumer Handler Started")
	select {} // Block forever
}

func (h *ReportHandlerImpl) startConsumer(queueName string, handler func([]byte) error) error {
	messages, err := h.rabbitMQ.ConsumeMessages(queueName)
	if err != nil {
		return fmt.Errorf("failed to consume messages from %s: %w", queueName, err)
	}

	go func() {
		for msg := range messages {
			logrus.Log.Infof("Message received from queue %s", queueName)
			if err := handler(msg.Body); err != nil {
				logrus.Log.Errorf("failed to handle message from %s: %v", queueName, err)
			}
		}
	}()

	return nil
}

func (h *ReportHandlerImpl) GetPriceHistoryReportFile(c *gin.Context) {
	commodityID, err := uuid.Parse(c.Param("commodity_id"))
	if err != nil {
		utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
		return
	}
	logrus.Log.Info("Commodity ID is valid ")

	cityID, err := strconv.ParseInt(c.Param("city_id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
		return
	}
	logrus.Log.Info("City ID is  valid")

	startDateStr := c.Query("start_date")
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
		return
	}
	endDatestr := c.Query("end_date")
	endDate, err := time.Parse("2006-01-02", endDatestr)
	if err != nil {
		utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
		return
	}

	params := &dto.PriceParamsDTO{
		CommodityID: commodityID,
		CityID:      cityID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	// Get the latest file (assuming filename contains timestamp)
	latestFile, err := h.excelUsecase.GetPriceExcelFile(c, params)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(*latestFile)))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Serve the file
	c.File(*latestFile)
}
func (h *ReportHandlerImpl) GetHarvestReportFile(c *gin.Context) {
	params := &dto.HarvestParamsDTO{}

	startDateStr := c.Query("start_date")
	if startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
			return
		}
		params.StartDate = startDate
	}
	endDatestr := c.Query("end_date")
	if endDatestr != "" {
		endDate, err := time.Parse("2006-01-02", endDatestr)
		if err != nil {
			utils.ErrorResponse(c, utils.NewBadRequestError(err.Error()))
			return
		}
		params.EndDate = endDate
	}

	id, err := uuid.Parse(c.Param("id"))
	params.LandCommodityID = id
	// Get the latest file (assuming filename contains timestamp)
	latestFile, err := h.excelUsecase.GetHarvestExcelFile(c, params)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(*latestFile)))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Serve the file
	c.File(*latestFile)
}
