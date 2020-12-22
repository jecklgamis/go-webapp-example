package server

import (
	"fmt"
	"github.com/gorilla/mux"
	handler "github.com/jecklgamis/go-api-server-template/pkg/server/handler"
	"log"
	"net/http"
	"time"
)

func Start(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/buildInfo", handler.BuildInfoHandler)
	router.HandleFunc("/probe/ready", handler.ReadinessProbeHandler)
	router.HandleFunc("/probe/live", handler.LivenessProbeHandler)
	router.HandleFunc("/api", handler.ApiHandler)
	router.HandleFunc("/", handler.RootHandler)

	addr := fmt.Sprintf("0.0.0.0:%d", port)

	fmt.Printf("Starting API server on  %s\n", addr)
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
