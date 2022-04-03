package trackingmetricstore

import (
	"sync"
	"testing"

	"MyMetric/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestTrackingMetrics_Add(t *testing.T) {
	cMetric := models.CustomMetric{
		Name:       "RandomValue",
		TypeM:      "gauge",
		UpdateFunc: func(args ...interface{}) (string, error) { return "", nil },
	}
	rMetric := models.RuntimeMetric{Name: "Alloc", TypeM: "gauge", UpdateFunc: func(args ...interface{}) (string, error) { return "", nil }}

	wantRuntimeM := []models.RuntimeMetric{rMetric}
	wantCcustomM := []models.CustomMetric{cMetric}

	t.Run("Test metric store", func(t *testing.T) {
		tr := &TrackingMetrics{
			mu:            &sync.Mutex{},
			RuntimeMetric: make([]models.RuntimeMetric, 0, 1),
			CustomMetric:  make([]models.CustomMetric, 0, 1),
		}
		tr.Add(cMetric)
		tr.Add(rMetric)

		actualMetricsC_ := tr.GetCustomMetrics()
		assert.Equal(t, wantCcustomM[0].Name, actualMetricsC_[0].Name, "GetCustomMetrics check Name")
		assert.Equal(t, wantCcustomM[0].TypeM, actualMetricsC_[0].TypeM, "GetCustomMetrics check TypeM")

		actualMetricsR_ := tr.GetRuntimeMetric()
		assert.Equal(t, wantRuntimeM[0].Name, actualMetricsR_[0].Name, "GetRuntimeMetric check Name")
		assert.Equal(t, wantRuntimeM[0].TypeM, actualMetricsR_[0].TypeM, "GetRuntimeMetric check TypeM")

	})
}
