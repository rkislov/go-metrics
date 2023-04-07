package entity

import (
	"errors"
	"github.com/google/uuid"
	"sync"
)

type Metric struct {
	ID    uuid.UUID
	Type  string
	Name  string
	Value float64
}

var (
	MetricsList = []Metric{
		{Type: "gauge", Name: "Alloc", Value: 0000},
		{Type: "gauge", Name: "BuckHashSys", Value: 0001},
	}
	ErrMissingValues = errors.New("пустые значения")
	ErrNotExists     = errors.New("такого ID нет")
)

type Storage interface {
	Add(m *Metric) error
	GetByID(id uuid.UUID) (Metric, error)
	GetAll() ([]Metric, error)
	Update(newMetric Metric) (Metric, error)
	GetByName(name string) (Metric, error)
}

type MemoryStorage struct {
	data map[uuid.UUID]Metric
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[uuid.UUID]Metric),
	}
}

func NewMetric(types, name string, value float64) (Metric, error) {
	if name == "" || types == "" {
		return Metric{}, ErrMissingValues
	}
	return Metric{
		ID:    uuid.New(),
		Type:  types,
		Name:  name,
		Value: value,
	}, nil
}

func (s *MemoryStorage) Add(m *Metric) error {
	s.Lock()
	defer s.Unlock()
	if m.Name == "" || m.Type == "" {
		return ErrMissingValues
	}
	s.data[m.ID] = *m

	return nil
}

func (s *MemoryStorage) GetByID(id uuid.UUID) (Metric, error) {
	s.Lock()
	defer s.Unlock()
	m, exists := s.data[id]
	if !exists {
		return Metric{}, ErrNotExists
	}

	return m, nil
}

func (s *MemoryStorage) GetAll() ([]Metric, error) {
	s.Lock()
	defer s.Unlock()
	var metrics []Metric

	for _, m := range s.data {
		metrics = append(metrics, m)
	}

	return metrics, nil
}

func (s *MemoryStorage) Update(newMetric Metric) (Metric, error) {
	s.Lock()
	defer s.Unlock()
	m, exists := s.data[newMetric.ID]
	if !exists {
		return Metric{}, ErrNotExists
	}
	s.data[m.ID] = newMetric

	return m, nil
}

func (s *MemoryStorage) GetByName(name string) (Metric, error) {
	s.Lock()
	defer s.Unlock()
	for _, m := range s.data {
		if m.Name == name {
			return m, nil
		}
	}
	return Metric{}, ErrNotExists
}
