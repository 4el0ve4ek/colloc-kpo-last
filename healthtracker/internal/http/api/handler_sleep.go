package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type sleepRequest struct {
	Duration int
}

func newSleepHandler(healthService healthService) *sleepHandler {
	return &sleepHandler{
		healthService: healthService,
	}
}

type sleepHandler struct {
	healthService healthService
}

func (h *sleepHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.handle(r)
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
}

func (h *sleepHandler) handle(r *http.Request) error {
	defer r.Body.Close()

	var request sleepRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return errors.Wrap(err, "decode request")
	}

	return errors.Wrap(h.healthService.AddSleep(request.Duration), "add sleep")
}
