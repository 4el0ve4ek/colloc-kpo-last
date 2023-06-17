package api

import "net/http"

func newSleepHandler() *sleepHandler {
	return &sleepHandler{}
}

type sleepHandler struct{}

func (h *sleepHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
