package routes

import (
	"net/http"
	"os"
	middleware "redis-cache-api/Helper/ApiHelper/Middleware"
	prom "redis-cache-api/Helper/ApiHelper/Prometheus"
	routeModel "redis-cache-api/Helper/ApiHelper/Routes/Models"
	subRoutes "redis-cache-api/Helper/ApiHelper/Routes/SubRoutes"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

var (
	appInfo = prom.SetGaugeMetric("redis_cache_api", "Application Info", "env", os.Getenv("env"), "version", os.Getenv("version"))
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

// NewRouter func() *Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Metrics endpoint for scrapping
	router.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(appInfo)
	router.Use(prometheusMiddleware)
	secureRouter := router.PathPrefix("/").Subrouter()
	// declare routes
	routes := []routeModel.Route{}
	routes = append(routes, subRoutes.GetAcceptTaskV1Routes()...)
	// bind routes
	for _, route := range routes {
		secureRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	secureRouter.Use(middleware.MiddlewareHandle)
	return router
}
