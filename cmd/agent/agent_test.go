package main

import (
	"context"
	"sync"
	"testing"

	"MyMetric/internal/agent"
	repositiry "MyMetric/internal/agent/repository"
	trackingMetricStore "MyMetric/internal/agent/tracking_metric_store"
	"MyMetric/internal/models"

	"github.com/stretchr/testify/assert"
)

type MockTransport struct {
	ch chan models.Metric
}

func (mt MockTransport) SendMetric(m models.Metric) error {
	mt.ch <- m
	return nil
}

func Test_agent(t *testing.T) {

	chForTest := make(chan models.Metric, 100)
	mockTransport := MockTransport{chForTest}

	listMetricForTrack := trackingMetricStore.New()
	listMetricForTrack.Add(models.CustomMetric{
		Name:  "PollCount",
		TypeM: "counter",
		UpdateFunc: func(args ...interface{}) (string, error) {
			return "10", nil
		}})

	a := agent.New(
		1,
		2,
		mockTransport,
		repositiry.NewRepoMem(),
		listMetricForTrack)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	go a.Start(ctx, cancel, wg)
	actualM := <-chForTest

	assert.Equal(t, "10", actualM.Value, "expect agent send 10 to chanel")
	cancel()
	close(chForTest)
	wg.Wait()
}
