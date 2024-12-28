package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
	"github.com/ryvasa/go-super-farmer-report-service/internal/model/dto"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	"github.com/ryvasa/go-super-farmer-report-service/utils"
	"github.com/xuri/excelize/v2"
)

type ExcelImpl struct {
	globFunc utils.GlobFunc
}

func NewExcelImpl(globFunc utils.GlobFunc) ExcelInterface {
	return &ExcelImpl{globFunc}
}

func (e *ExcelImpl) CreatePriceHistoryReport(results []domain.PriceHistory, commodityName, regionName string, commodityID uuid.UUID, cityID int64, startDate, endDate time.Time) error {
	// Buat file Excel baru
	f := excelize.NewFile()

	// Buat sheet baru
	sheetName := "Price History Report"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("error creating sheet: %v", err)
	}
	f.SetActiveSheet(index)

	// Set judul report
	f.SetCellValue(sheetName, "A1", fmt.Sprintf("Price History Report - %s in %s", commodityName, regionName))
	f.MergeCell(sheetName, "A1", "G1")

	// Tulis header
	headers := []string{"No", "Date", "Price", "Unit", "Commodity", "Region"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c3", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style untuk header
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return fmt.Errorf("error creating style: %v", err)
	}

	// Style untuk data numerik (rata kanan)
	numericStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Style untuk data teks (rata kiri)
	textStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Style untuk nomor urut (rata tengah)
	centerStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Apply style ke header
	f.SetCellStyle(sheetName, "A3", "F3", headerStyle)

	// Tulis data
	for i, record := range results {
		row := i + 4 // mulai dari baris ke-4
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i+1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.CreatedAt.Format("02-01-2006 15:04:05"))
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.Unit)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), commodityName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), regionName)

		// Apply style ke data
		f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), centerStyle)  // No (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), centerStyle)  // Date (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row), numericStyle) // Price (right)
		f.SetCellStyle(sheetName, fmt.Sprintf("D%d", row), fmt.Sprintf("D%d", row), centerStyle)  // Unit (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("E%d", row), textStyle)    // Commodity (left)
		f.SetCellStyle(sheetName, fmt.Sprintf("F%d", row), fmt.Sprintf("F%d", row), textStyle)    // Region (left)
	}

	// Auto-fit column width
	for i := 'A'; i <= 'F'; i++ {
		colName := string(i)
		width, _ := f.GetColWidth(sheetName, colName)
		if width < 15 {
			f.SetColWidth(sheetName, colName, colName, 15)
		}
	}

	// Buat nama file dengan timestamp
	fileName := fmt.Sprintf("./public/reports/price_history_%s_%d_%s_%s_%s.xlsx",
		commodityID,
		cityID,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
		time.Now().Format("20060102_150405"))

	// Simpan file
	if err := f.SaveAs(fileName); err != nil {
		return fmt.Errorf("error saving excel file: %v", err)
	}

	logrus.Log.WithField("Excel file created successfully:", fileName)
	return nil
}

func (e *ExcelImpl) CreateHarvestReport(results []domain.Harvest, commodityName, regionName, farmerName string, commodityID uuid.UUID, startDate, endDate time.Time) error {
	// Buat file Excel baru
	f := excelize.NewFile()

	// Buat sheet baru
	sheetName := "Harvest Report"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("error creating sheet: %v", err)
	}
	f.SetActiveSheet(index)

	// Set judul report
	f.SetCellValue(sheetName, "A1", fmt.Sprintf("Harvest Report - %s in %s", commodityName, regionName))
	f.MergeCell(sheetName, "A1", "G1")

	// Tulis header
	headers := []string{"No", "Date", "Quantity", "Unit", "Commodity", "Region", "Farmer"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c3", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style untuk header
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C6EFCE"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return fmt.Errorf("error creating style: %v", err)
	}

	// Style untuk data numerik (rata kanan)
	numericStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Style untuk data teks (rata kiri)
	textStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Style untuk nomor urut (rata tengah)
	centerStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	// Apply style ke header
	f.SetCellStyle(sheetName, "A3", "G3", headerStyle)

	// Tulis data
	for i, record := range results {
		row := i + 4 // mulai dari baris ke-4
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i+1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.CreatedAt.Format("02-01-2006 15:04:05"))
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.Unit)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), commodityName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), regionName)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), farmerName)

		// Style untuk data
		f.SetCellStyle(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), centerStyle)  // No (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), centerStyle)  // Date (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row), numericStyle) // Quantity (right)
		f.SetCellStyle(sheetName, fmt.Sprintf("D%d", row), fmt.Sprintf("D%d", row), centerStyle)  // Unit (center)
		f.SetCellStyle(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("E%d", row), textStyle)    // Commodity (left)
		f.SetCellStyle(sheetName, fmt.Sprintf("F%d", row), fmt.Sprintf("F%d", row), textStyle)    // Region (left)
		f.SetCellStyle(sheetName, fmt.Sprintf("G%d", row), fmt.Sprintf("G%d", row), textStyle)
	}

	// Auto-fit column width
	for i := 'A'; i <= 'G'; i++ {
		colName := string(i)
		width, _ := f.GetColWidth(sheetName, colName)
		if width < 15 {
			f.SetColWidth(sheetName, colName, colName, 15)
		}
	}

	// Buat nama file dengan timestamp
	fileName := fmt.Sprintf("./public/reports/harvests_%s_%s_%s_%s.xlsx",
		commodityID,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
		time.Now().Format("20060102_150405"))

	// Simpan file
	if err := f.SaveAs(fileName); err != nil {
		return fmt.Errorf("error saving excel file: %v", err)
	}

	logrus.Log.WithField("Excel file created successfully:", fileName)
	return nil
}

func (u *ExcelImpl) GetPriceExcelFile(ctx context.Context, params *dto.PriceParamsDTO) (*string, error) {
	// Get the latest excel file
	filePath := fmt.Sprintf("./public/reports/price_history_%s_%d_%s_%s_*.xlsx",
		params.CommodityID,
		params.CityID,
		params.StartDate.Format("2006-01-02"), params.EndDate.Format("2006-01-02"),
	)

	matches, err := u.globFunc.Glob(filePath) // Gunakan globFunc yang bisa dimock
	if err != nil {
		return nil, utils.NewInternalError("Error finding report file")
	}

	if len(matches) == 0 {
		return nil, utils.NewNotFoundError("Report file not found")
	}

	latestFile := matches[len(matches)-1]

	return &latestFile, nil
}

func (u *ExcelImpl) GetHarvestExcelFile(ctx context.Context, params *dto.HarvestParamsDTO) (*string, error) {
	filePath := fmt.Sprintf("./public/reports/harvests_%s_%s_%s_*.xlsx", params.LandCommodityID, params.StartDate.Format("2006-01-02"), params.EndDate.Format("2006-01-02"))
	matches, err := u.globFunc.Glob(filePath) // Gunakan globFunc yang bisa dimock
	if err != nil {
		return nil, utils.NewInternalError("Error finding report file")
	}

	if len(matches) == 0 {
		return nil, utils.NewNotFoundError("Report file not found")
	}

	// Get the latest file (assuming filename contains timestamp)
	latestFile := matches[len(matches)-1]

	return &latestFile, nil
}
