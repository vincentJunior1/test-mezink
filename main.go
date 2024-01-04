package main

import (
	"os"
	"os/signal"
	"skeleton-svc/controllers"
	v1Controllers "skeleton-svc/controllers/v1"
	"skeleton-svc/databases"
	"skeleton-svc/databases/postgre"
	"skeleton-svc/helpers"
	"skeleton-svc/routers"
	v1Usecases "skeleton-svc/usecases/v1"
	"syscall"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	helpers.LoadEnv()
	helpers.RegisterErrorCode()
	logger := helpers.InitializeNewLogs()
	// logger, gormLog := helpers.InitLogrus()

	psql := postgre.InitializePostgreDatabase(logger)

	db := databases.InitializeDatabase(
		psql,
		logger,
	)

	v1Usecase := v1Usecases.InitializeV1Usecase(
		db,
		logger,
	)
	v1Controller := v1Controllers.InitializeV1Controller(v1Usecase, logger)

	controller := controllers.InitializeController(
		v1Controller,
		logger,
	)
	router := routers.InitializeRouter(controller, logger)

	// sch := scheduler.InitializeScheduler(v1Usecase, logger)
	// sch.StartScheduler()

	logger.Info("Finish Initializing")

	// Start Server
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- router.StartServer()
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		logger.Info("got an interrupt, exiting...")
	case err := <-serverErr:
		if err != nil {
			logger.Error("error while running api, exiting...", err)
		}
	}

}
