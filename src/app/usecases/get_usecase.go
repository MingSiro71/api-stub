package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
	"api_stub/repositories"
)

type GetUsecase struct {
	repos map[string]interface{}
	out   outputs.GetOutput
}

func NewGetUsecase(repos map[string]interface{}, out outputs.GetOutput) *GetUsecase {
	return &GetUsecase{repos: repos, out: out}
}

func (uc *GetUsecase) Handle(dto dtos.QueryDto) error {
	id := dto.GetId()

	repo := uc.repos[repositories.Message].(repositories.MessageRepository)
	m, err := repo.Pop(id)

	if err != nil {
		return err
	}

	uc.out.Success(m)
	return nil
}
