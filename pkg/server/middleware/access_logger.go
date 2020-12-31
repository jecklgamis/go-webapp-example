package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// AccessLoggerMiddleware logs incoming HTTP request info
func AccessLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessLog := map[string]string{"host": r.Host,
			"method": r.Method, "uri_path": r.RequestURI, "protocol": r.Proto}
		bytes, _ := json.Marshal(accessLog)
		log.Println(string(bytes))
		next.ServeHTTP(w, r)
	})
}
