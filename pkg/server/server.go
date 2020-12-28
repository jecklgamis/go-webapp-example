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

func displayRoutes(router *mux.Router) {
	log.Println("Below are the configured endpoints")
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err == nil {
			log.Println(template)
		}
		return nil
	})
}

func Start() {
	env := GetEnvOrElse("APP_ENVIRONMENT", "dev")
	config := ReadConfig(env)
	router := mux.NewRouter()
	router.HandleFunc("/buildInfo", handler.BuildInfoHandler)
	router.HandleFunc("/probe/ready", handler.ReadinessProbeHandler)
	router.HandleFunc("/probe/live", handler.LivenessProbeHandler)
	router.HandleFunc("/api", handler.ApiHandler)
	router.HandleFunc("/", handler.RootHandler)
	router.Use(middleware.AccessLoggerMiddleware)
	displayRoutes(router)

	if config.Server.Https != nil {
		go func() {
			addr := fmt.Sprintf("0.0.0.0:%d", config.Server.Https.Port)
			log.Printf("Starting HTTPS server on %s\n", addr)
			srv := &http.Server{
				Handler:      router,
				Addr:         addr,
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}
			log.Fatal(srv.ListenAndServeTLS(config.Server.Https.CertFile, config.Server.Https.KeyFile))
		}()
	}
	if config.Server.Http != nil {
		go func() {
			addr := fmt.Sprintf("0.0.0.0:%d", config.Server.Http.Port)
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
