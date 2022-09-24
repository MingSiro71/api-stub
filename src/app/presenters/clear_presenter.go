package presenters

import (
	"api_stub/vo"
	"encoding/json"
	"fmt"
	"net/http"
)

type clearPresenter struct {
	w http.ResponseWriter
}

func NewClearPresenter(w http.ResponseWriter) *clearPresenter {
	return &clearPresenter{w: w}
}

func (p *clearPresenter) Success(id vo.Id, n int) {
	res := map[string]interface{}{"id": id.Tos(), "deletes": n}
	bytes, _ := json.Marshal(res)
	fmt.Fprint(p.w, string(bytes))
}

func (p *clearPresenter) Error(s string) {
	ShowError(p.w, s, 500)
}
