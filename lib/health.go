package random

import (
	"encoding/json"
	"net/http"
)

// NewHealthHandler is the constructor for HealthHandler.
func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		version: version,
	}
}

// HealthHandler provides a means to report on the health of the running web
// application. It reports the application's version.
type HealthHandler struct {
	version string
}

// ServeHTTP will return the applicaton's version as JSON.
func (h *HealthHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("content-type", "application/json; charset=utf-8")
	info := HealthInfo{Version: h.version}
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(info)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// HealthInfo represents a summary of the appliation's status.
type HealthInfo struct {
	Version string `json:"version"`
}
