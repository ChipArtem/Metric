package repository

import (
	"errors"
	"sync"

	"github.com/ChipArtem/Metric/internal/models"
)

type repoMem struct {
	mu          *sync.Mutex
	repoGauge   map[string]string
	repoCounter map[string]string
}

func NewRepoMem() *repoMem {
	return &repoMem{
		&sync.Mutex{},
		make(map[string]string),
		make(map[string]string),
	}
}

func (r *repoMem) GetAllMetric() []models.Metric {
	r.mu.Lock()
	defer r.mu.Unlock()
	m := make([]models.Metric, 0, 30)
	for name, value := range r.repoCounter {
		m = append(m, models.Metric{Name: name, Value: value, Type: "counter"})
	}
	for name, value := range r.repoGauge {
		m = append(m, models.Metric{Name: name, Value: value, Type: "gauge"})
	}
	return m
}

func (r *repoMem) GetMetric(tMetric, name string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	switch tMetric {
	case "gauge":
		v, ok := r.repoGauge[name]
		if !ok {
			return "", errors.New("not found")
		}
		return v, nil
	case "counter":
		v, ok := r.repoCounter[name]
		if !ok {
			return "", errors.New("not found")
		}
		return v, nil
	default:
		return "", errors.New("not found")
	}
}

func (r *repoMem) ExistMetric(tMetric, name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	switch tMetric {
	case "gauge":
		_, ok := r.repoGauge[name]
		return ok
	case "counter":
		_, ok := r.repoCounter[name]
		return ok
	}
	return false
}

func (r *repoMem) SetMetric(tMetric, name, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	switch tMetric {
	case "gauge":
		metrics := r.repoGauge
		metrics[name] = value
		r.repoGauge = metrics
	case "counter":
		metrics := r.repoCounter
		metrics[name] = value
		r.repoCounter = metrics
	}
}
