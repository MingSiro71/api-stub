package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
)

type GetUsecase struct {
	out outputs.GetOutput
}

func NewGetUsecase(out outputs.GetOutput) *GetUsecase {
	return &GetUsecase{out: out}
}

func (uc *GetUsecase) Handle(dto dtos.QueryDto) {
	// id := dto.GetId()

	// do something

	uc.out.Success("ok")
}
