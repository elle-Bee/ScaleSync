package metrics

import (
	"net/http"
	"time"
)

// MetricsMiddleware records HTTP metrics for each request.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Record incoming request
		ApiRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

		// Wrap response writer to capture status code
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		// Record the duration of the request
		duration := time.Since(start).Seconds()
		ApiRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		// Record success or failure
		if rec.statusCode >= 200 && rec.statusCode < 400 {
			ApiSuccesses.WithLabelValues(r.Method, r.URL.Path).Inc()
		} else {
			ApiFailures.WithLabelValues(r.Method, r.URL.Path).Inc()
		}
	})
}

// responseRecorder captures the HTTP status code for metrics.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
