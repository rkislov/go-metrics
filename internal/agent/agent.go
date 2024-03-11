package agent

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Server         string
	PollInterval   time.Duration
	ReportInterval time.Duration
}

const (
	gaugeTypeName   string = "gauge"
	counterTypeName string = "counter"
)

type CollectorAgent struct {
	cfg Config

	stats       runtime.MemStats
	PollCount   uint64
	RandomValue float64
}

func New(config Config) *CollectorAgent {
	collector := new(CollectorAgent)
	collector.cfg = config
	return collector
}

func (collector *CollectorAgent) CheckInit() (bool, error) {
	if collector.cfg.Server == "" {
		return false, errors.New("agent: Server field must be defined")
	}
	if collector.cfg.PollInterval == 0 {
		return false, errors.New("agent: PollInterval field must be defined")
	}
	if collector.cfg.ReportInterval == 0 {
		return false, errors.New("agent: ReportInterval field must be defined")
	}

	return true, nil
}

func (collector *CollectorAgent) Collect(t time.Time) {
	runtime.ReadMemStats(&collector.stats)
	collector.RandomValue = rand.Float64()
	collector.PollCount++
}

func (collector *CollectorAgent) PostOneGaugeStat(metricName string, metricValue float64) {
	url := "http://" + path.Join(collector.cfg.Server, "update", gaugeTypeName, metricName, strconv.FormatFloat(metricValue, 'f', -1, 64))
	resp, err := http.Post(url, "text/plain", strings.NewReader("body"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(url, " status code ", resp.StatusCode)
	}
	defer resp.Body.Close()
}

func (collector *CollectorAgent) PostOneCounterStat(metricName string, metricValue uint64) {
	url := "http://" + path.Join(collector.cfg.Server, "update", counterTypeName, metricName, strconv.FormatUint(metricValue, 10))
	resp, err := http.Post(url, "text/plain", strings.NewReader("body"))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(url, " status code ", resp.StatusCode)
	}
	defer resp.Body.Close()
}

func (collector *CollectorAgent) Report(t time.Time) {
	go collector.PostOneGaugeStat("Alloc", float64(collector.stats.Alloc))
	go collector.PostOneGaugeStat("BuckHashSys", float64(collector.stats.BuckHashSys))
	go collector.PostOneGaugeStat("Frees", float64(collector.stats.Frees))
	go collector.PostOneGaugeStat("GCCPUFraction", float64(collector.stats.GCCPUFraction))
	go collector.PostOneGaugeStat("GCSys", float64(collector.stats.GCSys))
	go collector.PostOneGaugeStat("HeapAlloc", float64(collector.stats.HeapAlloc))
	go collector.PostOneGaugeStat("HeapIdle", float64(collector.stats.HeapIdle))
	go collector.PostOneGaugeStat("HeapInuse", float64(collector.stats.HeapInuse))
	go collector.PostOneGaugeStat("HeapObjects", float64(collector.stats.HeapObjects))
	go collector.PostOneGaugeStat("HeapReleased", float64(collector.stats.HeapReleased))
	go collector.PostOneGaugeStat("HeapSys", float64(collector.stats.HeapSys))
	go collector.PostOneGaugeStat("LastGC", float64(collector.stats.LastGC))
	go collector.PostOneGaugeStat("Lookups", float64(collector.stats.Lookups))
	go collector.PostOneGaugeStat("MCacheInuse", float64(collector.stats.MCacheInuse))
	go collector.PostOneGaugeStat("MCacheSys", float64(collector.stats.MCacheSys))
	go collector.PostOneGaugeStat("MSpanInuse", float64(collector.stats.MSpanInuse))
	go collector.PostOneGaugeStat("MSpanSys", float64(collector.stats.MSpanSys))
	go collector.PostOneGaugeStat("Mallocs", float64(collector.stats.Mallocs))
	go collector.PostOneGaugeStat("NextGC", float64(collector.stats.NextGC))
	go collector.PostOneGaugeStat("NumForcedGC", float64(collector.stats.NumForcedGC))
	go collector.PostOneGaugeStat("NumGC", float64(collector.stats.NumGC))
	go collector.PostOneGaugeStat("OtherSys", float64(collector.stats.OtherSys))
	go collector.PostOneGaugeStat("PauseTotalNs", float64(collector.stats.PauseTotalNs))
	go collector.PostOneGaugeStat("StackInuse", float64(collector.stats.StackInuse))
	go collector.PostOneGaugeStat("StackSys", float64(collector.stats.StackSys))
	go collector.PostOneGaugeStat("Sys", float64(collector.stats.Sys))

	go collector.PostOneGaugeStat("RandomValue", float64(collector.RandomValue))

	go collector.PostOneCounterStat("PollCount", collector.PollCount)
}

func (collector *CollectorAgent) Run(end context.Context) error {
	ok, err := collector.CheckInit()
	if !ok {
		return err
	}

	collectTimer := time.NewTicker(collector.cfg.PollInterval)
	reportTimer := time.NewTicker(collector.cfg.ReportInterval)

	for {
		select {
		case t := <-collectTimer.C:
			collector.Collect(t)
		case t := <-reportTimer.C:
			collector.Report(t)
		case <-end.Done():
			return nil
		}
	}
}
