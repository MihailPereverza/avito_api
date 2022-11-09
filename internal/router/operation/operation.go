package operation

import (
	"avito_api/internal/handler/interface"
	"github.com/gorilla/mux"
)

func SetOperationRouter(accountHandler handler_interface.AccountHandlerInterface, router *mux.Router) {
	router.HandleFunc("/approve", accountHandler.ApproveDebiting).Methods("POST")
}
