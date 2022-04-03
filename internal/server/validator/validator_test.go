package validator

import (
	"testing"
)

func TestIsValidValue(t *testing.T) {
	type arg struct {
		mType string
		value string
	}
	tests := []struct {
		name string
		args arg
		want bool
	}{
		{"Positive gauge and float", arg{"gauge", "12.2"}, true},
		{"Negative gauge and string", arg{"gauge", "dsadsda"}, false},
		{"Positive gauge and float", arg{"gauge", "122"}, true},
		{"Negative gauge and string", arg{"gauge", "dsadsda"}, false},
		{"Negative wrong and float", arg{"gauge", "dsadsda"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidValue(tt.args.mType, tt.args.value); got != tt.want {
				t.Errorf("IsValidValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMType(t *testing.T) {

	tests := []struct {
		name  string
		mType string
		want  bool
	}{
		{"Positive 'gauge'", "gauge", true},
		{"Positive 'counter'", "counter", true},
		{"Negative 'NotVaildMType'", "notValid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMType(tt.mType); got != tt.want {
				t.Errorf("IsMType() = %v, want %v", got, tt.want)
			}
		})
	}
}
