package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/cmd/server/entity"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpMetricList []entity.Metric

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../../internal/templates/*")
	}

	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func SaveLists() {
	tmpMetricList = entity.MetricsList
}

func RestoreList() {
	entity.MetricsList = tmpMetricList
}
