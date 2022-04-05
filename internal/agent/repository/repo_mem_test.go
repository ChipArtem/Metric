package repositiry

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ChipArtem/Metric/internal/models"
)

func TestRepoMem_AddMetricValue(t *testing.T) {
	type mArg struct {
		mtype string
		name  string
		value string
	}

	tests := []struct {
		name string
		args mArg
	}{
		{name: "Add value in repomem",
			args: mArg{"type", "name", "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepoMem{
				mu:    &sync.Mutex{},
				store: make(map[string]string),
			}
			r.AddMetricValue(tt.args.mtype, tt.args.name, tt.args.value)
			for _, v := range r.store {
				assert.Equal(t, tt.args.value, v, "they should be equal")
			}
		})
	}
}

func TestRepoMem_GetMetricValue(t *testing.T) {
	type arg struct {
		mtype string
		name  string
	}
	tests := []struct {
		name    string
		args    arg
		want    string
		wantErr bool
	}{
		{name: "Get exist metric", args: arg{"mtype", "name"}, want: "value", wantErr: false},
		{name: "Get exist metric", args: arg{"notExist", "NotExistname"}, want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepoMem{
				mu:    &sync.Mutex{},
				store: make(map[string]string),
			}
			r.AddMetricValue("mtype", "name", "value")
			got, err := r.GetMetricValue(tt.args.mtype, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepoMem.GetMetricValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RepoMem.GetMetricValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepoMem_GetAll(t *testing.T) {
	test := struct {
		name string
		want []models.Metric
	}{name: "GetALl!", want: []models.Metric{{Type: "mtype", Name: "name", Value: "value"}}}
	t.Run(test.name, func(t *testing.T) {
		r := &RepoMem{
			mu:    &sync.Mutex{},
			store: make(map[string]string),
		}
		r.AddMetricValue("mtype", "name", "value")

		got, _ := r.GetAll()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RepoMem.GetAll() = %v, want %v", got, test.want)
		}
	})
}
