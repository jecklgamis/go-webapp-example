package server

import (
	"encoding/json"
	"net/http"
)

// RootHandler handles the /endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-webapp-example", "message": "It works on my machine!"})
}
