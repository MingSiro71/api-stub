package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
	"api_stub/repositories"
)

type ClearUsecase struct {
	repos map[string]interface{}
	out   outputs.ClearOutput
}

func NewClearUsecase(repos map[string]interface{}, out outputs.ClearOutput) *ClearUsecase {
	return &ClearUsecase{repos: repos, out: out}
}

func (uc *ClearUsecase) Handle(dto dtos.QueryDto) error {
	id := dto.GetId()

	repo := uc.repos[repositories.Message].(repositories.MessageRepository)
	l, err := repo.List(id)

	if err != nil {
		return err
	}

	err = repo.Clear(id)

	if err != nil {
		return err
	}

	uc.out.Success(id, len(l))
	return nil
}
