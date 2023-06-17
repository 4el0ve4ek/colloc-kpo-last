package api

import "net/http"

func newActivityHandler() *activityHandler {
	return &activityHandler{}
}

type activityHandler struct{}

func (h *activityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
