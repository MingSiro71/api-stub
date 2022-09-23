package controllers

import (
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func InitParam(r *http.Request) (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	vars := mux.Vars(r)

	if vars == nil {
		return params, errors.New("URI parse failed")
	}

	for k, v := range vars {
		params[k] = v
	}

	return params, nil
}

func ShowError(w http.ResponseWriter, m string) {
	e := map[string]string{
        "error": m,
    }
	bytes, _ := json.Marshal(e)
	j := string(bytes)

	fmt.Fprint(w, j)
}
