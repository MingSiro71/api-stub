package validations

import (
	"errors"
	"reflect"
)

func GetSureData(params map[string]interface{}) (string, error) {
	key := "data"
	v, exists := params[key]
	if exists != true {
		return "", errors.New("data is required")
	}

	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.String {
		return "", errors.New("data should string")
	}

	p := v.(string)
	return p, nil
}
