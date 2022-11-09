package app

import (
	"avito_api/internal/config"
	"context"
	"net/http"
	"time"
)

type App struct {
	cfg         *config.AppConfig
	router      http.Handler
	httpsServer *http.Server
}

func NewApp(cfg *config.AppConfig, router http.Handler) *App {
	return &App{
		cfg:    cfg,
		router: router,
		httpsServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        router,
			MaxHeaderBytes: 1 << 20, // 1MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (a *App) Start() error {
	return a.httpsServer.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.httpsServer.Shutdown(ctx)
}
