package config

import (
	"github.com/rkislov/go-metrics.git/internal/agent"
	"github.com/spf13/viper"
	"time"
)

func NewAgentConfig(v *viper.Viper) *agent.Config {
	v.SetDefault(envPollInterval, DefaultPollInterval)
	v.SetDefault(envReportInterval, DefaultReportInterval)
	v.SetDefault(envServer, DefaultServer)
	a := v.GetInt64(envPollInterval)
	println(a)
	return &agent.Config{
		PollInterval:   time.Duration(v.GetInt64(envPollInterval)) * time.Second,
		ReportInterval: time.Duration(v.GetInt64(envReportInterval)) * time.Second,
		Server:         v.GetString(envServer),
	}
}
