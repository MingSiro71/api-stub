package presenters

import (
	"encoding/json"
	"net/http"
)

func ShowError(w http.ResponseWriter, m string, c int) {
	e := map[string]string{
		"error": m,
	}
	bytes, _ := json.Marshal(e)
	http.Error(w, string(bytes), c)
}
