package vo

import (
	"errors"
	"regexp"
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
	res := regexp.MustCompile(`[a-zA-Z0-9]{6,256}`).Match([]byte(s))
	if res == false {
		return &id{value: ""}, errors.New(`id should 6-256 length with number and alphabet, recomended "com.company@username"`)
	}
	return &id{value: s}, nil
}

func (id *id) Tos() string {
	return id.value
}
