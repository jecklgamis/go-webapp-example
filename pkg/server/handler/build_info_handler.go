package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BuildInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-api-server-template",
		"version": "some-git-commit-id", "branch": "some-git-branch"})
	fmt.Fprintf(w, "{}")
}
