package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/internal/entity"
	"github.com/rkislov/go-metrics.git/internal/handlers"
	"log"
)

func main() {
	memoryStorage := entity.NewMemoryStorage()
	handler := handlers.NewHandler(memoryStorage)
	r := gin.Default()

	r.LoadHTMLGlob("../../internal/templates/*")

	r.GET("/", handler.ShowMetrics)
	r.POST("/update/:type/:name/:value", handler.UpdateOrCreate)

	log.Fatal(r.Run())
}
