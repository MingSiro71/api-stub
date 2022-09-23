package controllers

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v9"
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

func InitRedis(h string, p int, pw string, dbId int) *redis.Client {
	addr := strings.Join([]string{h, ":", strconv.Itoa(p)}, "")
 	redis := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: pw,
		DB: dbId,
	})

	return redis
}

func ShowError(w http.ResponseWriter, m string) {
	e := map[string]string{
        "error": m,
    }
	bytes, _ := json.Marshal(e)
	j := string(bytes)

	fmt.Fprint(w, j)
}
