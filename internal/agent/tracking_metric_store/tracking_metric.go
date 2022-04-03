package trackingmetricstore

import (
	"sync"

	"MyMetric/internal/models"
)

type TrackingMetrics struct {
	mu            *sync.Mutex
	RuntimeMetric []models.RuntimeMetric
	CustomMetric  []models.CustomMetric
}

func New() *TrackingMetrics {
	return &TrackingMetrics{
		mu:            &sync.Mutex{},
		RuntimeMetric: make([]models.RuntimeMetric, 0, 1),
		CustomMetric:  make([]models.CustomMetric, 0, 1),
	}
}

func (t *TrackingMetrics) Add(metrics ...interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	cM := t.CustomMetric
	rM := t.RuntimeMetric
	for _, m := range metrics {

		switch ms := m.(type) {
		case models.CustomMetric:
			cM = append(cM, ms)
		case models.RuntimeMetric:
			rM = append(rM, ms)
		case []models.RuntimeMetric:
			rM = append(rM, ms...)
		case []models.CustomMetric:
			cM = append(cM, ms...)
		}

	}
	t.CustomMetric = cM
	t.RuntimeMetric = rM
}

func (t *TrackingMetrics) GetCustomMetrics() []models.CustomMetric {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.CustomMetric
}

func (t *TrackingMetrics) GetRuntimeMetric() []models.RuntimeMetric {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.RuntimeMetric
}
