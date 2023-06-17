package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewServant(healthService healthService) {
	router := chi.NewRouter()

	router.Method(http.MethodPost, "/activity", newActivityHandler())
	router.Method(http.MethodPost, "/nutrition", newNutritionHandler())
	router.Method(http.MethodPost, "/activity", newSleepHandler())
	router.Method(http.MethodGet, "/stats", newStatsHandler())
}
