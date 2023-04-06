package main

import (
	"github.com/rkislov/go-metrics.git/cmd/server/entity"
	"github.com/rkislov/go-metrics.git/cmd/server/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowMetricsUnauthenticated(t *testing.T) {
	ms := entity.NewMemoryStorage()
	handler := handlers.NewHandler(ms)

	r := getRouter(true)

	r.GET("/", handler.ShowMetrics)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Metrics</title>") > 0

		return pageOK && statusOK
	})
}
