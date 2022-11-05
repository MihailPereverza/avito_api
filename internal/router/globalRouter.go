package router

import (
	"avito_api/internal/handler"
	"github.com/gorilla/mux"
)

type Router struct {
}

func (h *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	{
		user := api.PathPrefix("/user").Subrouter()
		{
			user.HandleFunc("/{id}", handler.Handler).Methods("GET")
		}
	}
	return router
}
