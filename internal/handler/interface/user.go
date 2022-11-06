package handler_interface

import (
	"net/http"
)

type UserHandlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
