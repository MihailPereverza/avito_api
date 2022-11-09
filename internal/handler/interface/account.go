package handler_interface

import (
	"net/http"
)

type AccountHandlerInterface interface {
	AddBalance(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
	GetStatistic(w http.ResponseWriter, r *http.Request)
	ReserveBalance(w http.ResponseWriter, r *http.Request)
	ApproveDebiting(w http.ResponseWriter, r *http.Request)
	UnReserveBalance(w http.ResponseWriter, r *http.Request)
	Transfer(w http.ResponseWriter, r *http.Request)
}
