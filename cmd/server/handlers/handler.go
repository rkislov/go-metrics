package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/cmd/server/entity"
	"net/http"
	"strconv"
)

type Handler struct {
	storage entity.Storage
}

func NewHandler(storage entity.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) ShowMetrics(c *gin.Context) {

	metrics, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Metrics",
			"payload": metrics,
		})
}

func (h *Handler) UpdateOrCreate(c *gin.Context) {
	var metric entity.Metric

	vp := c.Param("value")
	value, _ := strconv.ParseFloat(vp, 64)
	metric.Name = c.Param("name")
	metric.Type = c.Param("type")
	metric.Value = value

	existMetric, err := h.storage.GetByName(metric.Name)
	if err != nil {
		nm, err := entity.NewMetric(metric.Type, metric.Name, metric.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		}
		h.storage.Add(&nm)

	}
	metric.ID = existMetric.ID
	h.storage.Update(metric)
	fmt.Sprintf("%v", metric)

	c.JSON(http.StatusCreated, gin.H{"message": "ok"})

}
