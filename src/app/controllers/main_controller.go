package controllers

import (
	"context"
	"net/http"
	"io/ioutil"
	"api_stub/env"
	"api_stub/dtos"
	"api_stub/inputs"
	"api_stub/outputs"
	"api_stub/usecases"
	"api_stub/presenters"
	"api_stub/repositories"
	"api_stub/redis_repositories"
)

type MainController interface {
	Help(http.ResponseWriter, *http.Request)
	Set(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	// List(http.ResponseWriter, *http.Request)
	// Clear(http.ResponseWriter, *http.Request)
	// Init(http.ResponseWriter, *http.Request)
}
  
type mainController struct {  
}

func NewMainController() (m *mainController) {
	return &mainController{}
}

func (mc *mainController) Help(w http.ResponseWriter, r *http.Request) {
	out := presenters.NewHelpPresenter(w)
	out.Success()
}

func (mc *mainController) Set(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redis := InitRedis(env.RedisHost, env.RedisPort, env.RedisPassword, env.RedisDB)
	repos := map[string]interface{}{
		repositories.Message: redis_repositories.NewRedisMessageRepository(ctx, redis),
	}

	var out outputs.SetOutput = presenters.NewRegisterPresenter(w)
	var in inputs.SetInput = usecases.NewSetUsecase(repos, out)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ShowError(w, "parameters not validated.")
		return
	}

	params, err := InitParam(r)
	if err != nil {
		ShowError(w, "invalid url.")
		return
	}
	params["data"] = string(b)

	dto, err := dtos.NewSetDto(params)

	if err != nil {
		ShowError(w, err.Error())
		return
	}

	err = in.Handle(dto)
	if err != nil {
		ShowError(w, "internal server error.")
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
	dto, err := dtos.NewQueryDto(params)
	if err != nil {
		ShowError(w, err.Error())
		return
	}
	
	err = in.Handle(dto)
	if err != nil {
		ShowError(w, "internal server error.")
	}
}

// func (mc *mainController) List(w http.ResponseWriter, r *http.Request) {
// 	var out ListOutput
// 	out = NewListPresenter(w)
	
// 	var in ListInput
// 	in = NewMainUsecase(out)

// 	dto, err := NewListDto(params)
// 	if err != nil {
// 		w.write("parameters not validated")
// 	}
	
// 	in.List(dto)
// }

// func (mc *mainController) Clear(w http.ResponseWriter, r *http.Request) {
// 	var out ClearOutput
// 	out = NewClearPresenter(w)
	
// 	var in ClearInput
// 	in = NewMainUsecase(out)

// 	dto, err := NewClearDto(params)
// 	if err != nil {
// 		w.write("parameters not validated")
// 	}
	
// 	in.Clear(dto)
// }

// func (mc *mainController) Init(w http.ResponseWriter, r *http.Request) {
// 	var out InitOutput
// 	out = NewRegisterPresenter(w)
	
// 	var in InitInput
// 	in = NewMainUsecase(out)

// 	dto, err := NewInitrDto()
// 	// No paramater and validation 

// 	in.Init(dto)
// }
