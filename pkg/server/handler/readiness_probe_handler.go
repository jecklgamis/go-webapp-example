package server

import (
	"encoding/json"
	"net/http"
)

// ReadinessProbeHandler handles the /probe/ready endpoint
func ReadinessProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "I'm ready!"})
}
