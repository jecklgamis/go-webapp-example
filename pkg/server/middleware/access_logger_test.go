package server

import (
	test "github.com/jecklgamis/go-webapp-example/pkg/testing"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvokeNextHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	test.Assert(t, err == nil, "Unable to create request")

	nextHandlerInvoked := false
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextHandlerInvoked = true
	})
	rr := httptest.NewRecorder()
	handler := AccessLoggerMiddleware(nextHandler)
	handler.ServeHTTP(rr, req)
	test.Assert(t, nextHandlerInvoked, "Expecting next handler invocation")
}
