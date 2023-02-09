package metric

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

// Register TODO: fix dependency on httprouter
func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary heartbeat metric
// @Tags Metrics
// @Success 204
// @Failure 400
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
