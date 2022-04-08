package delaultmetrics

import (
	"math/rand"
	"reflect"
	"strconv"

	"github.com/ChipArtem/Metric/internal/agent/helpers"
	"github.com/ChipArtem/Metric/internal/models"
)

func funcUpdateRuntime(args ...interface{}) (string, error) {
	rValue := args[0].(reflect.Value)
	m := args[1].(models.RuntimeMetric)

	v := rValue.FieldByName(m.Name)
	value, err := helpers.RefValueToString(v, m.TypeM)
	return value, err
}

var (
	DefaultCustomMetric = []models.CustomMetric{
		{
			Name:  "PollCount",
			TypeM: "counter",
			UpdateFunc: func(args ...interface{}) (string, error) {
				s := args[0].(string)
				v64, err := strconv.ParseInt(s, 10, 64)
				v64 += 1
				return helpers.ValueToString(v64), err
			}},
		{
			Name:  "RandomValue",
			TypeM: "gauge",
			UpdateFunc: func(args ...interface{}) (string, error) {
				return helpers.ValueToString(rand.Float64()), nil
			},
		},
	}

	DefaultRuntimeMetric = []models.RuntimeMetric{
		{Name: "Alloc", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "BuckHashSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "Frees", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "GCCPUFraction", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "GCSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapAlloc", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapIdle", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapInuse", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapObjects", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapReleased", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "HeapSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "LastGC", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "Lookups", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "MCacheInuse", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "MCacheSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "MSpanInuse", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "MSpanSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "Mallocs", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "NextGC", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "NumForcedGC", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "NumGC", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "OtherSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "PauseTotalNs", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "StackInuse", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "StackSys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "Sys", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
		{Name: "TotalAlloc", TypeM: "gauge", UpdateFunc: funcUpdateRuntime},
	}
)
