package usecase

import (
	"strconv"

	"MyMetric/internal/models"
)

type MetricRepositorer interface {
	GetMetric(string, string) (string, error)
	SetMetric(string, string, string)
	ExistMetric(string, string) bool
	GetAllMetric() []models.Metric
}

type serverBusinessLogic struct {
	repo MetricRepositorer
}

func NewMetricBusinessLogic(r MetricRepositorer) *serverBusinessLogic {
	return &serverBusinessLogic{r}
}

func (bl *serverBusinessLogic) GetAll() []models.Metric {
	return bl.repo.GetAllMetric()
}

func (bl *serverBusinessLogic) GetMetric(mType, name string) (string, error) {
	v, err := bl.repo.GetMetric(mType, name)
	if err != nil {
		return "", err
	}
	return v, nil
}

func (bl *serverBusinessLogic) SetMetric(mType, name, value string) {
	if mType == "gauge" {
		bl.repo.SetMetric(mType, name, value)
	}
	if mType == "counter" {
		if bl.repo.ExistMetric(mType, name) {
			v, _ := bl.repo.GetMetric(mType, name)

			oldV, _ := strconv.ParseInt(v, 10, 64)
			newV, _ := strconv.ParseInt(value, 10, 64)
			newV += oldV
			value = strconv.FormatInt(newV, 10)

		}
		bl.repo.SetMetric(mType, name, value)
	}
}
