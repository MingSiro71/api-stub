package dtos

import (
	"api_stub/dtos/validations"
	"api_stub/vo"
)

type QueryDto interface {
	GetId() vo.Id
}

type queryDto struct {
	id vo.Id
}

func NewDummyQueryDto() QueryDto {
	return &queryDto{id: vo.NewDummyId()}
}

func NewQueryDto(params map[string]interface{}) (QueryDto, error) {
	id, err := validations.GetSureId(params)
	if err != nil {
		return NewDummyQueryDto(), err
	}

	return &queryDto{id: id}, nil
}

func (dto *queryDto) GetId() vo.Id {
	return dto.id
}
