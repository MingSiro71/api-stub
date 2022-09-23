package presenters

import (
	"fmt"
	"encoding/json"
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
