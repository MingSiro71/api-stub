package presenters

import (
	"api_stub/vo"
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

func (p *listPresenter) Success(id vo.Id, l []string) {
	s := make([]interface{}, 0)
	for _, v := range l {
		var m interface{}
		json.Unmarshal([]byte(v), &m)
		s = append(s, m)
	}

	d := map[string]interface{}{"id": id.Tos(), "messages": s}
	bytes, _ := json.Marshal(d)
	fmt.Fprint(p.w, string(bytes))
}
