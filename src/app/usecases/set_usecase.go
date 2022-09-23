package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
	"api_stub/repositories"
)

type SetUsecase struct {
	repos map[string]interface{}
	out outputs.SetOutput
}

func NewSetUsecase(repos map[string]interface{}, out outputs.SetOutput) *SetUsecase {
	return &SetUsecase{repos: repos, out: out}
}

func (uc *SetUsecase) Handle(dto dtos.SetDto) error {
	id := dto.GetId()
	data := dto.GetData()

	repo := uc.repos[repositories.Message].(repositories.MessageRepository)
	err := repo.Push(id, data)

	if err != nil {
		return err
	}

	uc.out.Success(id)
	return nil
}
