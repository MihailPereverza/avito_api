package handler_interface

import (
	"net/http"
)

type AccountHandlerInterface interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	AddBalance(w http.ResponseWriter, r *http.Request)
}
