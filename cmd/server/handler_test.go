package main

import (
	"github.com/rkislov/go-metrics.git/internal/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowMetricsUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handlers.ShowMetrics)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Metrics</title>") > 0

		return pageOK && statusOK
	})
}
