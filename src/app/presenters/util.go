package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ShowError(w http.ResponseWriter, m string) {
	e := map[string]string{
		"error": m,
	}
	bytes, _ := json.Marshal(e)
	j := string(bytes)

	fmt.Fprint(w, j)
}
