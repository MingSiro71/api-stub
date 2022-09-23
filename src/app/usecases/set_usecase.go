package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
)

type SetUsecase struct {
	out outputs.SetOutput
}

func NewSetUsecase(out outputs.SetOutput) *SetUsecase {
	return &SetUsecase{out: out}
}

func (uc *SetUsecase) Handle(dto dtos.SetDto) {
	id := dto.GetId()

	// do something

	uc.out.Success(id)
}
