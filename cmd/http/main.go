package main

import (
	"avito_api/internal/app"
	"avito_api/internal/config"
	"avito_api/internal/db"
	"avito_api/internal/router"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pg, err := db.NewPostgres(config.GetDBConfig())
	if err != nil {
		log.Fatal(err)
	}
	globalRouter := new(router.Router)
	srv := app.NewApp(config.GetAppConfig(), globalRouter.InitRoutes())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err = srv.Start(); err != nil {
			log.Printf("error occurred while http server was running: %s", err)
			quit <- syscall.SIGKILL
		}
	}()
	log.Println("server is running")

	<-quit
	log.Println("server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s\n", err.Error())
	}

	if err := pg.Close(); err != nil {
		log.Printf("error occured on pg connection close: %s", err.Error())
	}
	log.Println("server was shut down")
}
