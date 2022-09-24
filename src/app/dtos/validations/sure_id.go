package validations

import (
	"api_stub/exceptions"
	"api_stub/vo"
	"reflect"
)

const idKey = "id"

func GetSureId(params map[string]interface{}) (vo.Id, error) {
	v, exists := params[idKey]
	if exists != true {
		return vo.NewDummyId(), exceptions.NewValidationException(idKey + " is required")
	}

	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.String {
		return vo.NewDummyId(), exceptions.NewValidationException(idKey + "id should string")
	}

	p, err := vo.NewId(v.(string))
	return p, err
}
