package agent

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/ChipArtem/Metric/internal/models"
)

type MetricStorer interface {
	AddMetricValue(mtype, name, value string)
	GetMetricValue(mtype, name string) (string, error)
	GetAll() ([]models.Metric, error)
}

type Transporter interface {
	SendMetric(m models.Metric) error
}

type TrackingMetricsStorer interface {
	Add(...interface{})
	GetCustomMetrics() []models.CustomMetric
	GetRuntimeMetric() []models.RuntimeMetric
}

type Agent struct {
	pollInterval         int
	reportInterval       int
	trackingMetricsStore TrackingMetricsStorer
	metricStore          MetricStorer
	transport            Transporter
}

func New(
	secPollInterval int,
	secReportInterval int,
	client Transporter,
	metricStore MetricStorer,
	trackingMetricsStore TrackingMetricsStorer) *Agent {
	return &Agent{
		pollInterval:         secPollInterval,
		reportInterval:       secReportInterval,
		trackingMetricsStore: trackingMetricsStore,
		metricStore:          metricStore,
		transport:            client,
	}
}

func (a *Agent) AddMetrics(metrics ...interface{}) {
	a.trackingMetricsStore.Add(metrics)
}

func (a *Agent) updateMetric() {

	rms := &runtime.MemStats{}
	runtime.ReadMemStats(rms)
	r := reflect.Indirect(reflect.ValueOf(rms))

	for _, m := range a.trackingMetricsStore.GetRuntimeMetric() {
		v, _ := m.UpdateFunc(r, m)
		a.metricStore.AddMetricValue(m.TypeM, m.Name, v)
	}

	for _, m := range a.trackingMetricsStore.GetCustomMetrics() {
		v, _ := a.metricStore.GetMetricValue(m.TypeM, m.Name)
		value, _ := m.UpdateFunc(v)
		a.metricStore.AddMetricValue(m.TypeM, m.Name, value)
	}
}

func (a *Agent) sendMetric(ch chan<- models.Metric) {
	metrics, _ := a.metricStore.GetAll()
	for _, m := range metrics {
		ch <- m
	}
}

func (a *Agent) startSend(ctx context.Context, wg *sync.WaitGroup, ch <-chan models.Metric) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case v := <-ch:
			a.transport.SendMetric(v)
		}
	}
}

func (a *Agent) Start(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {

	mCh := make(chan models.Metric, 5)
	wg.Add(1)

	go a.startSend(ctx, wg, mCh)

	tPollInter := time.NewTicker(time.Duration(a.pollInterval) * time.Second)
	tRepInter := time.NewTicker(time.Duration(a.reportInterval) * time.Second)
	defer tPollInter.Stop()
	defer tRepInter.Stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case <-ctx.Done():
				close(mCh)
				break LOOP
			case <-tPollInter.C:
				a.updateMetric()
			case <-tRepInter.C:
				a.sendMetric(mCh)
			}
		}
	}()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

LOOP:
	for {
		s := <-signalChanel
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			cancel()
			break LOOP
		}
	}
	wg.Wait()
}
