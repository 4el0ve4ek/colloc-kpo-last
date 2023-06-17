package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"awesomeHealth/healthtracker/internal/http/api"
	"awesomeHealth/healthtracker/internal/repository/activity"
	"awesomeHealth/healthtracker/internal/repository/nutrition"
	"awesomeHealth/healthtracker/internal/repository/sleep"
	"awesomeHealth/healthtracker/internal/service"
	"awesomeHealth/healthtracker/pkg/database/postgres"
)

func main() {
	// HARDCODE
	dbConfig := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "username",
		Password: "password",
		DBName:   "health",
	}

	db, err := postgres.NewConn(dbConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "db conn setup"))
	}
	defer db.Close()

	activityRepository := activity.NewRepository(db)
	nutritionRepository := nutrition.NewRepository(db)
	sleepRepository := sleep.NewRepository(db)

	healthService := service.NewHealthService(activityRepository, nutritionRepository, sleepRepository)

	servant := api.NewServant(healthService)
	server := servant.GetServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	fmt.Println("listening on " + server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Print(errors.Wrap(err, "http server failure"))
			sigChan <- syscall.SIGINT
		}
	}()

	<-sigChan
	log.Println(errors.New("shutting down"))
}
