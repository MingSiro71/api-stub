package presenters

import (
	"fmt"
	"encoding/json"
	"net/http"
	"api_stub/vo"
)

type registerPresenter struct {
	w http.ResponseWriter
}

func NewRegisterPresenter(w http.ResponseWriter) (*registerPresenter) {
	return &registerPresenter{w: w}
}

func (p *registerPresenter) Success(id vo.Id) {
	res := make(map[string]string, 1)
	res["id"] = id.Tos()

	bytes, _:= json.Marshal(res)
	fmt.Fprint(p.w, string(bytes))
}

func (p *registerPresenter) Error(s string) {
	ShowError(p.w, s)
}
