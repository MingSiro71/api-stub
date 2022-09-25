package controllers

import (
	"api_stub/controllers/http_middlewares"
	"api_stub/dtos"
	"api_stub/env"
	"api_stub/exceptions"
	"api_stub/inputs"
	"api_stub/outputs"
	"api_stub/presenters"
	"api_stub/redis_repositories"
	"api_stub/repositories"
	"api_stub/usecases"
	"context"
	"io/ioutil"
	"net/http"
)

type MainController interface {
	Help(http.ResponseWriter, *http.Request)
	Set(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	Clear(http.ResponseWriter, *http.Request)
	Init(http.ResponseWriter, *http.Request)
}

type mainController struct {
	a http_middlewares.Auth
}

func NewMainController() (m *mainController) {
	return &mainController{a: http_middlewares.NewAuth()}
}

func (mc *mainController) Set(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.SetOutput = presenters.NewSetPresenter(w)
	var in inputs.SetInput = usecases.NewSetUsecase(repos, out)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ShowError(w, exceptions.NewValidationException(exceptions.ValidationExceptionDefault))
		return
	}

	params, err := InitParam(r)
	if err != nil {
		ShowError(w, exceptions.NewRoutingException(exceptions.RoutingExceptionDefault))
		return
	}
	params["data"] = string(b)

	dto, err := dtos.NewSetDto(params)
	if err != nil {
		ShowError(w, err)
		return
	}

	err = in.Handle(dto)
	if err != nil {
		ShowError(w, err)
	}
}

func (mc *mainController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.GetOutput = presenters.NewGetPresenter(w)
	var in inputs.GetInput = usecases.NewGetUsecase(repos, out)

	params, err := InitParam(r)
	if err != nil {
		ShowError(w, exceptions.NewRoutingException(exceptions.RoutingExceptionDefault))
		return
	}

	dto, err := dtos.NewQueryDto(params)
	if err != nil {
		ShowError(w, err)
		return
	}

	err = in.Handle(dto)
	if err != nil {
		ShowError(w, err)
	}
}

func (mc *mainController) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.ListOutput = presenters.NewListPresenter(w)
	var in inputs.ListInput = usecases.NewListUsecase(repos, out)

	params, err := InitParam(r)
	if err != nil {
		ShowError(w, exceptions.NewRoutingException(exceptions.RoutingExceptionDefault))
		return
	}

	dto, err := dtos.NewQueryDto(params)
	if err != nil {
		ShowError(w, err)
		return
	}

	err = in.Handle(dto)
	if err != nil {
		ShowError(w, err)
	}
}

func (mc *mainController) Clear(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.ClearOutput = presenters.NewClearPresenter(w)
	var in inputs.ClearInput = usecases.NewClearUsecase(repos, out)

	params, err := InitParam(r)
	if err != nil {
		ShowError(w, exceptions.NewRoutingException(exceptions.RoutingExceptionDefault))
		return
	}

	dto, err := dtos.NewQueryDto(params)
	if err != nil {
		ShowError(w, err)
		return
	}

	err = in.Handle(dto)
	if err != nil {
		ShowError(w, err)
	}
}

func (mc *mainController) Init(w http.ResponseWriter, r *http.Request) {
	if mc.a.IsAdmin(r) == false {
		ShowError(w, exceptions.NewAuthException(exceptions.AuthExceptionDefault))
		return
	}
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.InitOutput = presenters.NewInitPresenter(w)
	var in inputs.InitInput = usecases.NewInitUsecase(repos, out)

	err := in.Handle()
	if err != nil {
		ShowError(w, err)
		return
	}
}
