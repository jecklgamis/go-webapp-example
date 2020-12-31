package server

import (
	"fmt"
	"github.com/gorilla/mux"
	handler "github.com/jecklgamis/go-api-server-template/pkg/server/handler"
	middleware "github.com/jecklgamis/go-api-server-template/pkg/server/middleware"
	"github.com/jecklgamis/go-api-server-template/pkg/version"
	"log"
	"net/http"
	"time"
)

func printRoutes(router *mux.Router) {
	log.Println("Below are the configured endpoints")
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		if err == nil {
			log.Printf("%s %s\n", methods, template)
		}
		return nil
	})
}

// Start starts the HTTP server
func Start() {
	env := GetEnvOrElse("APP_ENVIRONMENT", "dev")
	config := ReadConfig(env)
	router := mux.NewRouter()
	router.HandleFunc("/buildInfo", handler.BuildInfoHandler).Methods(http.MethodGet)
	router.HandleFunc("/probe/ready", handler.ReadinessProbeHandler).Methods(http.MethodGet)
	router.HandleFunc("/probe/live", handler.LivenessProbeHandler).Methods(http.MethodGet)
	router.HandleFunc("/api", handler.APIHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/", handler.RootHandler).Methods(http.MethodGet)
	router.Use(middleware.AccessLoggerMiddleware)
	printRoutes(router)

	if config.Server.HTTPS != nil {
		go func() {
			addr := fmt.Sprintf("0.0.0.0:%d", config.Server.HTTPS.Port)
			log.Printf("Starting HTTPS server on %s\n", addr)
			srv := &http.Server{
				Handler:      router,
				Addr:         addr,
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}
			log.Fatal(srv.ListenAndServeTLS(config.Server.HTTPS.CertFile, config.Server.HTTPS.KeyFile))
		}()
	}
	if config.Server.HTTP != nil {
		go func() {
			addr := fmt.Sprintf("0.0.0.0:%d", config.Server.HTTP.Port)
			log.Printf("Starting HTTP server on %s\n", addr)
			srv := &http.Server{
				Handler:      router,
				Addr:         addr,
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}
			log.Fatal(srv.ListenAndServe())
			log.Printf("Version: %s\n", version.BuildVersion)
		}()
	}
	for {
		time.Sleep(time.Second)
	}
}
