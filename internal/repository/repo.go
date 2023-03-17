package repository

import "github.com/rkislov/go-metrics.git/internal/entity"

type MetricRepository interface {
	//Migrate() error
	AddMetric(metric entity.Metric) []entity.Metric
	Update(metric entity.Metric) bool
	GetAll() []entity.Metric
	IsExist(name string) bool
}
