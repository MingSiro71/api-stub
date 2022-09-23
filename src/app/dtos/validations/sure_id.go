package validations

import (
	"api_stub/vo"
	"errors"
	"reflect"
)

func GetSureId(params map[string]interface{}) (vo.Id, error) {
	key := "id"
	v, exists := params[key]
	if exists != true {
		return vo.NewDummyId(), errors.New("id is required")
	}

	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.String {
		return vo.NewDummyId(), errors.New("id should string")
	}

	p, err := vo.NewId(v.(string))
	return p, err
}
