package account

import (
	"avito_api/internal/handler/interface"
	"github.com/gorilla/mux"
)

func SetAccountRouter(accountHandler handler_interface.AccountHandlerInterface, router *mux.Router) {
	router.HandleFunc("/", accountHandler.CreateAccount).Methods("POST")
	router.HandleFunc("/balance/add", accountHandler.AddBalance).Methods("POST")
}
