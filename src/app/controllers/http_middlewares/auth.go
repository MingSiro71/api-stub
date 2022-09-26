package http_middlewares

import (
	"net/http"
	"os"
)

const adminKeyHeader = "X-Admin-Key"

type Auth interface {
	IsAdmin(*http.Request) bool
}

type auth struct {
}

func NewAuth() *auth {
	return &auth{}
}

func (a *auth) IsAdmin(r *http.Request) bool {
	if r.Header[adminKeyHeader][0] == string(os.Getenv("ADMIN_KEY")) {
		return true
	}
	return false
}
