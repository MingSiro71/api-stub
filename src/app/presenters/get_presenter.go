package presenters

import (
	"fmt"
	"net/http"
)

type getPresenter struct {
	w http.ResponseWriter
}

func NewGetPresenter(w http.ResponseWriter) *getPresenter {
	return &getPresenter{w: w}
}

func (p *getPresenter) Success(j string) {
	fmt.Fprint(p.w, j)
}
