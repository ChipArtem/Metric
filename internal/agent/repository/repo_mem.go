package repositiry

import (
	"strings"
	"sync"

	"github.com/ChipArtem/Metric/internal/models"
)

type RepoMem struct {
	mu    sync.Mutex
	store map[string]string
}

func NewRepoMem() *RepoMem {
	return &RepoMem{mu: sync.Mutex{}, store: make(map[string]string)}
}

func (r *RepoMem) AddMetricValue(mtype, name, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	m := r.store
	m[mtype+"#@#@"+name] = value
	r.store = m
}

func (r *RepoMem) GetMetricValue(mtype, name string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, err := r.store[mtype+"#@#@"+name]
	if !err {
		return "", nil
	}
	return v, nil
}

func (r *RepoMem) GetAll() ([]models.Metric, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	newMp := make([]models.Metric, 0, len(r.store))
	for k, v := range r.store {
		typeName := strings.Split(k, "#@#@")
		newMp = append(
			newMp,
			models.Metric{Type: typeName[0], Name: typeName[1], Value: v},
		)
	}
	return newMp, nil
}
