package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type activityRequest struct {
	Activity string
	Duration int
	Calories int
}

func newActivityHandler(healthService healthService) *activityHandler {
	return &activityHandler{
		healthService: healthService,
	}
}

type activityHandler struct {
	healthService healthService
}

func (h *activityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *activityHandler) handle(r *http.Request) error {
	defer r.Body.Close()

	var request activityRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return errors.Wrap(err, "decode request")
	}

	return errors.Wrap(h.healthService.AddActivity(request.Activity, request.Duration, request.Calories), "add activity")
}
