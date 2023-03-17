package repository

import "github.com/rkislov/go-metrics.git/internal/entity"

type inMemRepo struct {
	data map[entity.Metric]*entity.Metric
}

func NewInMemRepo() MetricRepository {
	return inMemRepo{
		data: map[entity.Metric]*entity.Metric{},
	}
}

func (r inMemRepo) GetAll() []entity.Metric {
	return entity.MetricsList
}

func (r inMemRepo) AddMetric(metric entity.Metric) []entity.Metric {
	return append(entity.MetricsList, metric)
}
func (r inMemRepo) Update(metric entity.Metric) bool {
	for i, _ := range entity.MetricsList {
		if entity.MetricsList[i].Type == metric.Type &&
			entity.MetricsList[i].Name == metric.Name {
			entity.MetricsList[i].Value = metric.Value
		}
		return false
	}
	return true
}

func (r inMemRepo) IsExist(name string) bool {
	for i, _ := range entity.MetricsList {
		if entity.MetricsList[i].Type == name {
			return true
		}

	}
	return false
}
