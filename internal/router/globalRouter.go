package router

import (
	"avito_api/internal/handler/interface"
	account2 "avito_api/internal/router/account"
	user2 "avito_api/internal/router/user"
	"github.com/gorilla/mux"
)

type Router struct {
	UserHandler    handler_interface.UserHandlerInterface
	AccountHandler handler_interface.AccountHandlerInterface
	GlobalRouter   *mux.Router
}

func NewRouter(
	userHandler handler_interface.UserHandlerInterface,
	accountHandler handler_interface.AccountHandlerInterface,
) *Router {
	return &Router{
		UserHandler:    userHandler,
		AccountHandler: accountHandler,
		GlobalRouter:   mux.NewRouter(),
	}
}

func (r *Router) InitRoutes() *mux.Router {
	api := r.GlobalRouter.PathPrefix("/api").Subrouter()
	{
		user := api.PathPrefix("/user").Subrouter()
		{
			user2.SetUserRouter(r.UserHandler, user)
		}
		account := api.PathPrefix("/account").Subrouter()
		{
			account2.SetAccountRouter(r.AccountHandler, account)
		}
	}
	return r.GlobalRouter
}
