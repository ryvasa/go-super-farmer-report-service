package report_handler

type Handlers struct {
	ReportHandler ReportHandler
}

func NewHandlers(reportHandler ReportHandler) *Handlers {
	return &Handlers{
		ReportHandler: reportHandler,
	}
}
