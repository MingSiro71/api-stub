package usecases

import (
	"api_stub/outputs"
	"api_stub/repositories"
)

type InitUsecase struct {
	repos map[string]interface{}
	out   outputs.InitOutput
}

func NewInitUsecase(repos map[string]interface{}, out outputs.InitOutput) *InitUsecase {
	return &InitUsecase{repos: repos, out: out}
}

func (uc *InitUsecase) Handle() error {
	repo := uc.repos[repositories.Message].(repositories.MessageRepository)
	err := repo.Init()

	if err != nil {
		return err
	}

	uc.out.Success()
	return nil
}
