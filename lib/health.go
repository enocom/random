package random

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// NewHealthHandler is the constructor for HealthHandler.
func NewHealthHandler(version string) *HealthHandler {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate Health guid: %v", err)
	}
	guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return &HealthHandler{
		guid:    guid,
		version: version,
	}
}

// HealthHandler provides a means to report on the health of the running web
// application. It reports the application's version.
type HealthHandler struct {
	guid    string
	version string
}

// ServeHTTP will return the applicaton's version as JSON.
func (h *HealthHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("content-type", "application/json; charset=utf-8")
	info := HealthInfo{
		GUID:    h.guid,
		Version: h.version,
	}
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(info)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// HealthInfo represents a summary of the appliation's status.
type HealthInfo struct {
	GUID    string `json:"guid"`
	Version string `json:"version"`
}
