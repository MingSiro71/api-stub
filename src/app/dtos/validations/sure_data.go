package validations

import (
	"api_stub/exceptions"
	"reflect"
)

const dataKey = "data"

func GetSureData(params map[string]interface{}) (string, error) {
	v, exists := params[dataKey]
	if exists != true {
		return "", exceptions.NewValidationException(dataKey + " is required")
	}

	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.String {
		return "", exceptions.NewValidationException(dataKey + " should string")
	}

	p := v.(string)
	return p, nil
}
