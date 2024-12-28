package usecase

type ReportUsecase interface {
	HandlePriceHistoryMessage(msgBody []byte) error
	HandleHarvestMessage(msgBody []byte) error
}
