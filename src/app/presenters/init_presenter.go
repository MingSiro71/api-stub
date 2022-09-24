package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type initPresenter struct {
	w http.ResponseWriter
}

func NewInitPresenter(w http.ResponseWriter) *initPresenter {
	return &initPresenter{w: w}
}

func (p *initPresenter) Success() {
	res := map[string]string{"result": "successed"}
	bytes, _ := json.Marshal(res)
	fmt.Fprint(p.w, string(bytes))
}

func (p *initPresenter) Error(s string) {
	ShowError(p.w, s, 500)
}
