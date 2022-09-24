package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type listPresenter struct {
	w http.ResponseWriter
}

func NewListPresenter(w http.ResponseWriter) *listPresenter {
	return &listPresenter{w: w}
}

func (p *listPresenter) Success(l []string) {
	var s []interface{}
	for _, v := range l {
		var m interface{}
		json.Unmarshal([]byte(v), &m)
		s = append(s, m)
	}

	bytes, _ := json.Marshal(s)
	fmt.Fprint(p.w, string(bytes))
}

func (p *listPresenter) Error(s string) {
	ShowError(p.w, s)
}
