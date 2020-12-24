package server

import (
	"encoding/json"
	"github.com/jecklgamis/go-api-server-template/pkg/version"
	"net/http"
)

func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-api-server-template",
		"version": version.BuildVersion, "branch": version.BuildBranch})
}
