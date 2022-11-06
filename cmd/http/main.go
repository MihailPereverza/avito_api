package main

import (
	"avito_api/internal/app"
	"avito_api/internal/config"
	"avito_api/internal/db"
	"avito_api/internal/db/account"
	user2 "avito_api/internal/db/user"
	account2 "avito_api/internal/handler/account"
	"avito_api/internal/handler/user"
	"avito_api/internal/router"
	"avito_api/internal/service"
	"context"
	"database/sql"
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
	globalRouter := createRouter(pg)
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

func createRouter(db *sql.DB) *router.Router {
	userDB := user2.NewDBUser(db)
	userService := service.NewUserService(userDB)
	userHandler := user.NewUserHandler(userService)

	accountDB := account.NewDBAccount(db)
	accountService := service.NewAccountService(accountDB)
	accountHandler := account2.NewAccountHandler(accountService)

	//currencyDB := currency.NewDBCurrency(db)
	//fmt.Println(currencyDB.GetAllCurrency())

	return router.NewRouter(userHandler, accountHandler)
}
