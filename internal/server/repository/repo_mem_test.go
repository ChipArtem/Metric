package repository

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repoMem_SetMetric(t *testing.T) {
	repoGauge := make(map[string]string)
	repoCounter := make(map[string]string)

	type Metric struct {
		tMetric string
		name    string
		value   string
	}
	tests := []struct {
		name string
		args Metric
	}{
		{"Set gauge metric", Metric{"gauge", "name", "value"}},
		{"Set counter metric", Metric{"counter", "name", "value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMem{
				mu:          &sync.Mutex{},
				repoGauge:   repoGauge,
				repoCounter: repoCounter,
			}
			r.SetMetric(tt.args.tMetric, tt.args.name, tt.args.value)
			if tt.args.tMetric == "gauge" {
				assert.Equal(t, tt.args.value, repoGauge[tt.args.name])
			}
			if tt.args.tMetric == "counter" {
				assert.Equal(t, tt.args.value, repoCounter[tt.args.name])
			}
		})
	}
}

func Test_repoMem_ExistMetric(t *testing.T) {

	repoGauge := map[string]string{"name": "value"}
	repoCounter := map[string]string{"name": "value"}

	type arg struct {
		tMetric string
		name    string
	}
	tests := []struct {
		name string
		args arg
		want bool
	}{
		{"Exist gauge metric", arg{"gauge", "name"}, true},
		{"Not exist gauge metric", arg{"gauge", "NotExist"}, false},
		{"Exist counter metric", arg{"counter", "name"}, true},
		{"Not exist counter metric", arg{"counter", "NotExist"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMem{
				mu:          &sync.Mutex{},
				repoGauge:   repoGauge,
				repoCounter: repoCounter,
			}
			if got := r.ExistMetric(tt.args.tMetric, tt.args.name); got != tt.want {
				t.Errorf("repoMem.ExistMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoMem_GetMetric(t *testing.T) {
	repoGauge := map[string]string{"name": "value"}
	repoCounter := map[string]string{"name": "value"}

	type arg struct {
		tMetric string
		name    string
	}
	tests := []struct {
		name    string
		args    arg
		want    string
		wantErr bool
	}{
		{"Exist gauge metric", arg{"gauge", "name"}, "value", false},
		{"Not exist gauge metric", arg{"gauge", "NotExist"}, "", true},
		{"Exist counter metric", arg{"counter", "name"}, "value", false},
		{"Not exist counter metric", arg{"counter", "NotExist"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMem{
				mu:          &sync.Mutex{},
				repoGauge:   repoGauge,
				repoCounter: repoCounter,
			}
			got, err := r.GetMetric(tt.args.tMetric, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoMem.GetMetric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repoMem.GetMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoMem_GetAllMetric(t *testing.T) {

	repoGauge := map[string]string{"name": "value"}
	repoCounter := map[string]string{"name": "value"}

	t.Run("", func(t *testing.T) {
		r := &repoMem{
			mu:          &sync.Mutex{},
			repoGauge:   repoGauge,
			repoCounter: repoCounter,
		}
		got := r.GetAllMetric()
		assert.Equal(t, 2, len(got), "expect repoMem.GetAllMetric() return 2")
	})
}
