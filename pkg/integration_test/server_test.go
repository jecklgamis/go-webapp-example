package server_test

import (
	"fmt"
	it "github.com/jecklgamis/go-api-server-template/pkg/integration_test"
	"github.com/jecklgamis/go-api-server-template/pkg/server"
	test "github.com/jecklgamis/go-api-server-template/pkg/testing"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"testing"
)

func TestServerEndPoints(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "dev")
	port := it.UnusedPort()
	if testing.Short() {
		t.Skip()
	}
	go func() {
		viper.Set("SERVER.HTTP.PORT", fmt.Sprintf("%d", port))
		server.Start()
	}()
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	r, err := http.Get(fmt.Sprintf("%s/buildInfo", baseURL))
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /buildInfo")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /buildInfo")

	r, _ = http.Get(fmt.Sprintf("%s/probe/ready", baseURL))
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /probe/ready")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /probe/ready")

	r, _ = http.Get(fmt.Sprintf("%s/probe/live", baseURL))
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /probe/live")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /probe/live")

	r, _ = http.Get(fmt.Sprintf("%s/api", baseURL))
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /api")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /api")

	r, _ = http.Get(fmt.Sprintf("%s/metrics", baseURL))
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /metrics")

	r, _ = http.Get(baseURL)
	test.Assert(t, err == nil, "Unable to send request")
	test.Assert(t, r.StatusCode == http.StatusOK, "Unable to reach /")
	test.Assert(t, r.Header.Get("Content-Type") == "application/json", "Unexpected Content-Type from /")

}
