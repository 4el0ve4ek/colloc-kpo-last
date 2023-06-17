package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type nutritionRequest struct {
	Dish     string
	Size     int
	Calories int
}

func newNutritionHandler(healthService healthService) *nutritionHandler {
	return &nutritionHandler{
		healthService: healthService,
	}
}

type nutritionHandler struct {
	healthService healthService
}

func (h *nutritionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *nutritionHandler) handle(r *http.Request) error {
	defer r.Body.Close()

	var request nutritionRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return errors.Wrap(err, "decode request")
	}

	return errors.Wrap(h.healthService.AddNutrition(request.Dish, request.Size, request.Calories), "add nutrition")
}
