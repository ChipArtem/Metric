package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/ChipArtem/Metric/internal/agent"
	delaultMetrics "github.com/ChipArtem/Metric/internal/agent/delault_metrics"
	repositiry "github.com/ChipArtem/Metric/internal/agent/repository"
	trackingMetricStore "github.com/ChipArtem/Metric/internal/agent/tracking_metric_store"
	"github.com/ChipArtem/Metric/internal/agent/transport"
)

func main() {

	listMetricForTrack := trackingMetricStore.New()
	listMetricForTrack.Add(delaultMetrics.DefaultRuntimeMetric)
	listMetricForTrack.Add(delaultMetrics.DefaultCustomMetric)

	a := agent.New(
		2,
		10,
		transport.NewHTTPClient("127.0.0.1:8080", &http.Client{}),
		repositiry.NewRepoMem(),
		listMetricForTrack)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	a.Start(ctx, cancel, wg)
}
