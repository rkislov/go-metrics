package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/internal/entity"
	"github.com/rkislov/go-metrics.git/internal/repository"
	"net/http"
	"strconv"
)

func ShowMetrics(c *gin.Context) {
	repo := repository.NewInMemRepo()
	metrics := repo.GetAll()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Metrics",
			"payload": metrics,
		})
}

func UpdateOrCreate(c *gin.Context) {
	repo := repository.NewInMemRepo()
	var metric entity.Metric
	vp := c.Param("value")
	value, _ := strconv.ParseFloat(vp, 64)
	metric.Name = c.Param("name")
	metric.Type = c.Param("type")
	metric.Value = value
	if repo.IsExist(metric.Name) {
		repo.Update(metric)
	} else {
		repo.AddMetric(metric)
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}
