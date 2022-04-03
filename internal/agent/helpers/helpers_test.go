package helpers

import (
	"reflect"
	"testing"
)

func Test_getTypeValue(t *testing.T) {

	var TestUint64 uint64 = 184467440
	var TestFloat64 = 102.9
	valueUint64 := reflect.Indirect(reflect.ValueOf(TestUint64))
	valueFloat64 := reflect.Indirect(reflect.ValueOf(TestFloat64))

	type arg struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args arg
		want string
	}{
		{"Test uint64", arg{valueUint64}, "uint64"},
		{"Test float64", arg{valueFloat64}, "float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTypeValue(tt.args.v); got != tt.want {
				t.Errorf("getTypeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefValueToString(t *testing.T) {

	var TestUint64 uint64 = 184467440
	var TestFloat64 = 102.9

	valueUint64 := reflect.Indirect(reflect.ValueOf(TestUint64))
	valueFloat64 := reflect.Indirect(reflect.ValueOf(TestFloat64))

	type arg struct {
		v     reflect.Value
		mType string
	}

	tests := []struct {
		name    string
		args    arg
		want    string
		wantErr bool
	}{
		{"mType gauge uint64", arg{valueUint64, "gauge"}, "184467440.000000", false},
		{"mType gauge float 64", arg{valueFloat64, "gauge"}, "102.900000", false},
		{"mType counter uint64", arg{valueUint64, "counter"}, "184467440", false},
		{"mType counter float64", arg{valueFloat64, "counter"}, "102", false},
		{"Wrong type ", arg{valueFloat64, "wrongType"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RefValueToString(tt.args.v, tt.args.mType)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("%s, RefValueToString() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestValueToString(t *testing.T) {
	type arg struct {
		v interface{}
	}
	tests := []struct {
		name string
		args arg
		want string
	}{
		{"Convert int64", arg{int64(-9223372036854775808)}, "-9223372036854775808"},
		{"Convert uint64", arg{uint64(18446744073709551615)}, "18446744073709551615"},
		{"Convert float64", arg{float64(100100.100001)}, "100100.100001"},
		{"Convert complex128", arg{complex128(10)}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueToString(tt.args.v); got != tt.want {
				t.Errorf("ValueToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
