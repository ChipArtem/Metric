package models

type Metric struct {
	Name  string
	Type  string
	Value string
}

type RuntimeMetric struct {
	Name       string
	TypeM      string
	UpdateFunc func(...interface{}) (string, error)
}

type CustomMetric struct {
	Name       string
	TypeM      string
	UpdateFunc func(...interface{}) (string, error)
}
