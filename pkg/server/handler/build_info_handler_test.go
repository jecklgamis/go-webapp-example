package server

import (
	"encoding/json"
	"fmt"
	test "github.com/jecklgamis/go-api-server-template/pkg/testing"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/buildInfo", nil)
	test.Assert(t, err == nil, "Unable to create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BuildInfoHandler)
	handler.ServeHTTP(rr, req)
	test.Assert(t, rr.Code == http.StatusOK, "Unexpected status code")
	test.Assert(t, rr.Header().Get("Content-Type") == "application/json",
		fmt.Sprintf("Unexpected content type %s", rr.Header().Get("Content-Type")))

	var entity map[string]string
	json.Unmarshal(rr.Body.Bytes(), &entity)
	test.Assert(t, entity["name"] == "go-api-server-template", "Unexpected name")
	test.Assert(t, entity["version"] == "", "Unexpected version")
	test.Assert(t, entity["branch"] == "", "Unexpected branch")
}
