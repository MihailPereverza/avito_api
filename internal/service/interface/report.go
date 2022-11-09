package service_interface

type ReportServiceInterface interface {
	GetReport() (string, error)
}
