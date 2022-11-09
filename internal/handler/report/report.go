package hanlderReport

import (
	handler_interface "avito_api/internal/handler/interface"
	"avito_api/internal/service/interface"
)

type ReportHandler struct {
	reportService service_interface.ReportServiceInterface
}

func NewReportHandler(
	reportService service_interface.ReportServiceInterface,
) handler_interface.ReportHandlerInterface {
	return &ReportHandler{
		reportService: reportService,
	}
}
