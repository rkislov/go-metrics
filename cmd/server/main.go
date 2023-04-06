package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/cmd/server/entity"
	"github.com/rkislov/go-metrics.git/cmd/server/handlers"
	"log"
)

func main() {
	memoryStorage := entity.NewMemoryStorage()
	handler := handlers.NewHandler(memoryStorage)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", handler.ShowMetrics)
	r.POST("/update/:type/:name/:value", handler.UpdateOrCreate)

	log.Fatal(r.Run())
}
