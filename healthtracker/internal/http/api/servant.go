package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func NewServant(healthService healthService) *servant {
	router := chi.NewRouter()

	router.Method(http.MethodPost, "/activity", newActivityHandler(healthService))
	router.Method(http.MethodPost, "/nutrition", newNutritionHandler(healthService))
	router.Method(http.MethodPost, "/sleep", newSleepHandler(healthService))
	router.Method(http.MethodGet, "/stats", newStatsHandler(healthService))

	server := &http.Server{
		ReadHeaderTimeout: time.Minute,
		Addr:              fmt.Sprintf(":%d", 8080), // HARDCODE
		Handler:           router,
	}
	return &servant{server: server}
}

type servant struct {
	server *http.Server
}

func (s *servant) GetServer() *http.Server {
	return s.server
}
