package presenters

import (
	"api_stub/vo"
	"encoding/json"
	"fmt"
	"net/http"
)

type setPresenter struct {
	w http.ResponseWriter
}

func NewSetPresenter(w http.ResponseWriter) *setPresenter {
	return &setPresenter{w: w}
}

func (p *setPresenter) Success(id vo.Id) {
	res := map[string]string{"id": id.Tos()}
	bytes, _ := json.Marshal(res)
	fmt.Fprint(p.w, string(bytes))
}

func (p *setPresenter) Error(s string) {
	ShowError(p.w, s)
}
