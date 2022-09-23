package dtos

import (
	"api_stub/vo"
	"api_stub/dtos/validations"
)

type SetDto interface {
	GetId() vo.Id
	GetData() string
}

type setDto struct {
	id vo.Id
	data string
}

func NewDummySetDto() SetDto {
	return &setDto{id: vo.NewDummyId(), data: ""}
}

func NewSetDto(params map[string]interface{}) (SetDto, error) {
	id, err := validations.GetSureId(params)
	if err != nil {
		return NewDummySetDto(), err
	}

	data, err := validations.GetSureData(params)
	if err != nil {
		return NewDummySetDto(), err
	}

	return &setDto{id: id, data: data}, nil
}

func (dto *setDto) GetId() vo.Id {
	return dto.id
}

func (dto *setDto) GetData() string {
	return dto.data
}
