package account

import (
	"avito_api/internal/handler/interface"
	"github.com/gorilla/mux"
)

func SetAccountRouter(accountHandler handler_interface.AccountHandlerInterface, router *mux.Router) {
	router.HandleFunc("/balance/add", accountHandler.AddBalance).Methods("POST")
	router.HandleFunc("/balance/reserve", accountHandler.ReserveBalance).Methods("POST")
	router.HandleFunc("/balance/unreserve", accountHandler.UnReserveBalance).Methods("POST")
	router.HandleFunc("/balance", accountHandler.GetBalance).Methods("POST")
	router.HandleFunc("/{accountID}/statistic", accountHandler.GetStatistic).Methods("GET")
	router.HandleFunc("/transfer", accountHandler.Transfer).Methods("POST")
}
