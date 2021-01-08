package server

import (
	"fmt"
	"github.com/gorilla/mux"
	handler "github.com/jecklgamis/go-api-server-template/pkg/server/handler"
	"github.com/jecklgamis/go-api-server-template/pkg/server/middleware"
	"github.com/jecklgamis/go-api-server-template/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func printRoutes(router *mux.Router) {
	log.Println("Available endpoints:")
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		if err == nil {
			log.Printf("%s %s\n", methods, template)
		}
		return nil
	})
}

var (
	httpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of http requests",
	}, []string{"code", "method"})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of all HTTP requests",
	}, []string{"code", "handler", "method"})
)

func instrumentHandler(f http.HandlerFunc, label string) http.Handler {
	return promhttp.InstrumentHandlerDuration(
		httpRequestDuration.MustCurryWith(prometheus.Labels{"handler": label}),
		promhttp.InstrumentHandlerCounter(httpRequestTotal, f))

}

// Start starts the HTTP server
func Start() {
	env := GetEnvOrElse("APP_ENVIRONMENT", "dev")
	config := ReadConfig(env)
	router := mux.NewRouter()

	prometheus.DefaultRegisterer.MustRegister(httpRequestTotal)
	prometheus.DefaultRegisterer.MustRegister(httpRequestDuration)

	router.Handle("/buildInfo", instrumentHandler(handler.BuildInfoHandler, "build_info")).Methods(http.MethodGet)
	router.Handle("/probe/ready", instrumentHandler(handler.ReadinessProbeHandler, "ready")).Methods(http.MethodGet)
	router.Handle("/probe/live", instrumentHandler(handler.LivenessProbeHandler, "live")).Methods(http.MethodGet)
	router.Handle("/api", instrumentHandler(handler.APIHandler, "api")).Methods(http.MethodGet, http.MethodPost)
	router.Handle("/", instrumentHandler(handler.RootHandler, "api")).Methods(http.MethodGet)
	router.Handle("/metrics", promhttp.Handler())
	router.Use(server.AccessLoggerMiddleware)

	printRoutes(router)

	if config.Server != nil && config.Server.HTTPS != nil {
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
	if config.Server != nil && config.Server.HTTP != nil {
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
