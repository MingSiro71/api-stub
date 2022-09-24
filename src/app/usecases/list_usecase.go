package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
	"api_stub/repositories"
)

type ListUsecase struct {
	repos map[string]interface{}
	out   outputs.ListOutput
}

func NewListUsecase(repos map[string]interface{}, out outputs.ListOutput) *ListUsecase {
	return &ListUsecase{repos: repos, out: out}
}

func (uc *ListUsecase) Handle(dto dtos.QueryDto) error {
	id := dto.GetId()

	repo := uc.repos[repositories.Message].(repositories.MessageRepository)
	l, err := repo.List(id)

	if err != nil {
		return err
	}

	uc.out.Success(id, l)
	return nil
}
