package validator

import "github.com/asaskevich/govalidator"

func IsValidValue(mType, value string) bool {
	if mType == "gauge" {
		return govalidator.IsFloat(value)
	}
	if mType == "counter" {
		return govalidator.IsInt(value)
	}
	return false
}
func IsMType(mType string) bool {
	if mType == "gauge" || mType == "counter" {
		return true
	}
	return false
}
