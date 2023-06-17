package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func newStatsHandler(healthService healthService) *statsHandler {
	return &statsHandler{
		healthService: healthService,
	}
}

type statsHandler struct {
	healthService healthService
}

func (h *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := h.handle(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte(err.Error()))
		if errWrite != nil {
			log.Print(errWrite)
		}
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(resp)
	if errWrite != nil {
		log.Print(errWrite)
	}
}

func (h *statsHandler) handle(r *http.Request) ([]byte, error) {
	stats, err := h.healthService.GetStats()
	if err != nil {
		return nil, errors.Wrap(err, "get stats")
	}
	resp, err := json.Marshal(stats)
	if err != nil {
		return nil, errors.Wrap(err, "marshal json")
	}
	return resp, nil
}
