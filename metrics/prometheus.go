package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsServer() {
	registerRequestMetrics(prometheus.DefaultRegisterer)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    "localhost:9394",
		Handler: mux,
	}

	go func() {
		log.Println("starting metrics server ...")
		log.Println(server.ListenAndServe())
	}()
}
