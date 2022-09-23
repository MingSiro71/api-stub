package usecases

import (
	"api_stub/dtos"
	"api_stub/outputs"
	"api_stub/repositories"
	"github.com/go-redis/redis/v9"
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

	if err == redis.Nil {
		uc.out.Error("no data stored.")
	} else if err != nil {
		return err
	}

	uc.out.Success(m)
	return nil
}
