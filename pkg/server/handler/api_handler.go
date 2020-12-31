package server

import (
	"encoding/json"
	"net/http"
)

// APIHandler handles the /api endpoint
func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-api-server-template", "message": "You have reached the /api endpoint!"})
}
