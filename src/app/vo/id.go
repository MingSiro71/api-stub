package vo

import (
	"errors"
)

type Id interface {
	Tos() string
}

type id struct {
	value string
}

func NewDummyId() Id {
	return &id{value: ""}
}

func NewId(s string) (Id, error) {
	// validate s
	if false {
		return &id{value: ""}, errors.New("invalid paramater: id")
	}
	return &id{value: s}, nil
}

func (id *id) Tos() string {
	return id.value
}