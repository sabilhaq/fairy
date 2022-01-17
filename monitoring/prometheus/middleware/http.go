// Package middleware is a prometheus middleware for HTTP server.
package middleware

import (
	"net/http"
	"strconv"
	"time"

	mwChi "github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

const httpReqName = "http_requests_total"
const httpLatencyName = "http_request_duration_seconds"

var hm httpMiddleware

type httpMiddleware struct {
	req *prometheus.CounterVec
	lat *prometheus.HistogramVec
}

func init() {
	hm = httpMiddleware{
		req: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: httpReqName,
				Help: "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			},
			[]string{"code", "method", "path"},
		),
		lat: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: httpLatencyName,
				Help: "How long it took to process the request, partitioned by status code, method and HTTP path.",
			},
			[]string{"code", "method", "path"},
		),
	}

	prometheus.MustRegister(hm.req)
	prometheus.MustRegister(hm.lat)
}

// NewHTTP to create new prometheus http middleware.
func NewHTTP() func(http.Handler) http.Handler {
	return hm.handler
}

func (mw httpMiddleware) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := mwChi.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		mw.req.WithLabelValues(strconv.Itoa(ww.Status()), r.Method, r.URL.Path).Inc()
		mw.lat.WithLabelValues(strconv.Itoa(ww.Status()), r.Method, r.URL.Path).Observe(float64(time.Since(start).Seconds()))
	})
}