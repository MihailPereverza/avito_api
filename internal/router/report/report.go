package report

import (
	"avito_api/internal/handler/interface"
	"github.com/gorilla/mux"
)

func SetReportRouter(reportHandler handler_interface.ReportHandlerInterface, router *mux.Router) {
	router.HandleFunc("/", reportHandler.GetReport).Methods("GET")
}
