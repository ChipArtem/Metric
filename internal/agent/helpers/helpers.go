package helpers

import (
	"fmt"
	"reflect"
	"strconv"
)

func ValueToString(v interface{}) string {
	switch m := v.(type) {
	case float64:
		return fmt.Sprintf("%f", m)
	case int64:
		return strconv.FormatInt(m, 10)
	case uint64:
		return strconv.FormatUint(m, 10)
	}
	return ""
}

func RefValueToString(v reflect.Value, mType string) (string, error) {
	vType := getTypeValue(v)
	switch vType {
	case "float64":
		if mType == "gauge" {
			return ValueToString(v.Float()), nil
		}
		if mType == "counter" {
			return ValueToString(int64(v.Float())), nil
		}

	case "uint64", "uint32":
		if mType == "gauge" {
			return ValueToString(float64(v.Uint())), nil
		}
		if mType == "counter" {
			return ValueToString(int64(v.Uint())), nil
		}
	}
	return "", fmt.Errorf("not supported type %s", vType)
}

func getTypeValue(v reflect.Value) string {
	return v.Type().String()
}
