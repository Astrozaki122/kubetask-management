package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kubetask_http_requests_total",
			Help: "Total HTTP requests to KubeTask service",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequests)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpRequests.WithLabelValues("/").Inc()
		fmt.Fprintln(w, "KubeTask Management System")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		httpRequests.WithLabelValues("/health").Inc()
		w.Write([]byte("ok"))
	})

	// 🔥 PROMETHEUS METRICS ENDPOINT
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
