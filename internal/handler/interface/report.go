package handler_interface

import "net/http"

type ReportHandlerInterface interface {
	GetReport(w http.ResponseWriter, r *http.Request)
}
