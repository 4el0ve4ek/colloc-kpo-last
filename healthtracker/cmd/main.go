package cmd

import (
	"log"

	"github.com/pkg/errors"

	"awesomeHealth/healthtracker/internal/repository/activity"
	"awesomeHealth/healthtracker/internal/repository/nutrition"
	"awesomeHealth/healthtracker/internal/repository/sleep"
	"awesomeHealth/healthtracker/internal/service"
	"awesomeHealth/healthtracker/pkg/database/postgres"
)

func main() {

	// HARDCODE
	dbConfig := postgres.Config{
		Host:     "postgres",
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

}
