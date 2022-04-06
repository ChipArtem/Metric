package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ChipArtem/Metric/internal/agent"
	repositiry "github.com/ChipArtem/Metric/internal/agent/repository"
	trackingMetricStore "github.com/ChipArtem/Metric/internal/agent/tracking_metric_store"
	transport "github.com/ChipArtem/Metric/internal/agent/transport"
	"github.com/ChipArtem/Metric/internal/models"
)

type MockTransport struct {
	ch chan models.Metric
}

func (mt MockTransport) SendMetric(m models.Metric) error {
	mt.ch <- m
	return nil
}

func Test_agent(t *testing.T) {

	ch := make(chan string, 100)
	ctx, cancel := context.WithCancel(context.Background())

	svr := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ch <- r.URL.Path
			}))
	defer svr.Close()

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
		transport.NewHTTPClient(svr.URL, &http.Client{}),
		repositiry.NewRepoMem(),
		listMetricForTrack)

	wg := &sync.WaitGroup{}

	go a.Start(ctx, cancel, wg)

	actualValue := <-ch
	fmt.Println("END")
	assert.Equal(t, "/update/counter/PollCount/10", actualValue, "expect agent send 10 to chanel")
	cancel()
	svr.Close()
	close(ch)
}
