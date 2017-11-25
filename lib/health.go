package random

import (
	"encoding/json"
	"net/http"
)

func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		version: version,
	}
}

type HealthHandler struct {
	version string
}

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

type HealthInfo struct {
	Version string `json:"version"`
}
