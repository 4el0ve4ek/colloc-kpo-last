package api

import "net/http"

func newStatsHandler() *statsHandler {
	return &statsHandler{}
}

type statsHandler struct{}

func (h *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
