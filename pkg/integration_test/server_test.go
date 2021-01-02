package server_test

import (
	"github.com/jecklgamis/go-api-server-template/pkg/server"
	test "github.com/jecklgamis/go-api-server-template/pkg/testing"
	"net/http"
	"os"
	"testing"
)

func TestServerEndPoints(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "dev")
	go func() {
		server.Start()
	}()
	r, _ := http.Get("http://localhost:8080/buildInfo")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /buildInfo")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /buildInfo")

	r, _ = http.Get("http://localhost:8080/probe/ready")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /probe/ready")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /probe/ready")

	r, _ = http.Get("http://localhost:8080/probe/live")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /probe/live")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /probe/live")

	r, _ = http.Get("http://localhost:8080/api")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /api")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /api")

	r, _ = http.Get("http://localhost:8080/")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /")
}
