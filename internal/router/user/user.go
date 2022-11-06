package user

import (
	"avito_api/internal/handler/interface"
	"github.com/gorilla/mux"
)

func SetUserRouter(userHandler handler_interface.UserHandlerInterface, router *mux.Router) {
	router.HandleFunc("/", userHandler.CreateUser).Methods("POST")
}
