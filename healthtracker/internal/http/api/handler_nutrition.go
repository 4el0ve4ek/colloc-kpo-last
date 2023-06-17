package api

import "net/http"

func newNutritionHandler() *nutritionHandler {
	return &nutritionHandler{}
}

type nutritionHandler struct{}

func (h *nutritionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
