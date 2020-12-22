package server

import (
	"encoding/json"
	"net/http"
)

func LivenessProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "I'm alive!"})
}
