package server

import (
	"encoding/json"
	"github.com/jecklgamis/go-webapp-example/pkg/version"
	"net/http"
)

// BuildInfoHandler handles the /buildInfo endpoint
func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache;no-store;max-age=0")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-webapp-example",
		"version": version.BuildVersion, "branch": version.BuildBranch})
}
