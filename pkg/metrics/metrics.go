package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Metrics for tracking HTTP requests, errors, and durations
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests handled by the application",
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets, // Default buckets for histogram
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	HttpRequestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of errors encountered in HTTP requests",
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)
	// Metrics for tracking API requests, errors, and durations by endpoint
	ApiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests handled by the application",
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	ApiRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Duration of API requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	ApiFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_failures_total",
			Help: "Total number of failed API requests",
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	ApiSuccesses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_successes_total",
			Help: "Total number of successful API requests",
		},
		[]string{"endpoint"}, // Labels: endpoint name
	)

	// Metrics for tracking user creation (new users)
	NewUserCreated = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "new_user_created_total",
			Help: "Total number of new users created",
		},
		[]string{"method"}, // Labels: method (POST)
	)

	// Metrics for tracking user login attempts
	UserLoginAttempts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_login_attempts_total",
			Help: "Total number of user login attempts",
		},
		[]string{"method"}, // Labels: method (POST, GET)
	)
)

// InitMetrics registers all the metrics with Prometheus.
func InitMetrics() {
	// Register all metrics
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(HttpRequestErrors)
	prometheus.MustRegister(ApiRequests)
	prometheus.MustRegister(ApiRequestDuration)
	prometheus.MustRegister(ApiFailures)
	prometheus.MustRegister(ApiSuccesses)
	prometheus.MustRegister(NewUserCreated)
	prometheus.MustRegister(UserLoginAttempts)
}

// MetricsHandler returns the handler for Prometheus metrics scraping.
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

// RegisterMetrics registers custom metrics with Prometheus
func RegisterMetrics() {
	prometheus.MustRegister(ApiRequests, ApiSuccesses, ApiFailures, ApiRequestDuration)
}
