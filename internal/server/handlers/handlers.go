package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"github.com/ChipArtem/Metric/internal/models"
	"github.com/ChipArtem/Metric/internal/server/validator"
)

type MetricBusinessLogicer interface {
	SetMetric(string, string, string)
	GetAll() []models.Metric
	GetMetric(string, string) (string, error)
}

type metricHandler struct {
	bl   MetricBusinessLogicer
	host string
}

func NewMetricHandler(bl MetricBusinessLogicer, host string) metricHandler {
	return metricHandler{bl, host}
}

func (m *metricHandler) GetMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mType := vars["mtype"]
	name := vars["name"]

	v, err := m.bl.GetMetric(mType, name)
	if err != nil {
		http.Error(w, ``, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(v)); err != nil {
		fmt.Printf("\n(m *metricHandler) SetMetric %s", err)
	}
}

func (m *metricHandler) SetMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mType := vars["mtype"]
	name := vars["name"]
	value := vars["value"]
	if !validator.IsValidValue(mType, value) {
		http.Error(w, ``, http.StatusBadRequest)
		return
	}
	m.bl.SetMetric(mType, name, value)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(``)); err != nil {
		fmt.Printf("\n(m *metricHandler) SetMetric %s", err)
	}
}

func (m *metricHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "Metric")[0]
	tmpl := template.Must(template.ParseFiles(path + "/Metric/internal/template/index.html"))

	metrics := m.bl.GetAll()
	if err := tmpl.Execute(w, struct{ Metrics []models.Metric }{metrics}); err != nil {
		fmt.Printf("\n(m *metricHandler) GetAll %s", err)
	}
}
