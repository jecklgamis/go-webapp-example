package server

import (
	"encoding/json"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": "go-api-server-template", "message": "You have reached the /api endpoint!"})
}
