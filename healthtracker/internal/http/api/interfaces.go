package api

import "awesomeHealth/healthtracker/internal/models"

type healthService interface {
	AddNutrition(dish string, size int, calories int) error
	AddSleep(durationInSeconds int) error
	AddActivity(activity string, durationInSeconds int, calories int) error
	GetStats() (models.Stats, error)
}
