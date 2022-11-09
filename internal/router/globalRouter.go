package router

import (
	"avito_api/internal/handler/interface"
	account2 "avito_api/internal/router/account"
	operation2 "avito_api/internal/router/operation"
	report2 "avito_api/internal/router/report"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"

	_ "avito_api/docs"
)

type Router struct {
	AccountHandler handler_interface.AccountHandlerInterface
	GlobalRouter   *mux.Router
	ReportHandler  handler_interface.ReportHandlerInterface
}

func NewRouter(
	accountHandler handler_interface.AccountHandlerInterface,
	reportHandler handler_interface.ReportHandlerInterface,
) *Router {
	return &Router{
		AccountHandler: accountHandler,
		ReportHandler:  reportHandler,
		GlobalRouter:   mux.NewRouter(),
	}
}

func (r *Router) InitRoutes() *mux.Router {
	r.GlobalRouter.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	api := r.GlobalRouter.PathPrefix("/api").Subrouter()
	{
		account := api.PathPrefix("/account").Subrouter()
		{
			account2.SetAccountRouter(r.AccountHandler, account)
		}
		operation := api.PathPrefix("/operation").Subrouter()
		{
			operation2.SetOperationRouter(r.AccountHandler, operation)
		}
		report := api.PathPrefix("/report").Subrouter()
		{
			report2.SetReportRouter(r.ReportHandler, report)
		}
	}

	// доступ к статическим файлам отчетов в csv
	r.GlobalRouter.PathPrefix("/reports/").Handler(
		http.StripPrefix(
			"/reports/",
			http.FileServer(http.Dir("./internal/static/reports/")),
		),
	).Methods("GET")
	return r.GlobalRouter
}
