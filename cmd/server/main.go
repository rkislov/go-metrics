package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/go-metrics.git/internal/handlers"
	"log"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("../../internal/templates/*")

	r.GET("/", handlers.ShowMetrics)
	r.POST("/update/:type/:name/:value", handlers.UpdateOrCreate)

	log.Fatal(r.Run())
}
