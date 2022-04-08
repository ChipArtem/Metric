package handlers

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/ChipArtem/Metric/internal/models"
	"github.com/ChipArtem/Metric/internal/server/validator"

	"github.com/gorilla/mux"
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
	w.Write([]byte(v))

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
	w.Write([]byte(``))
}

func (m *metricHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "Metric")[0]
	tmpl := template.Must(template.ParseFiles(path + "/Metric/internal/template/index.html"))

	metrics := m.bl.GetAll()
	tmpl.Execute(w, struct{ Metrics []models.Metric }{metrics})
}
