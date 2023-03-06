package main

import (
	"github.com/rkislov/go-metrics.git/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/update/gauge/", handlers.GaugeHandler)
	http.HandleFunc("/update/counter/", handlers.CounterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
