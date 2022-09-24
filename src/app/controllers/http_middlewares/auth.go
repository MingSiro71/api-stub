package http_middlewares

import (
	"api_stub/env"
	"net/http"
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
	if r.Header[adminKeyHeader][0] == string(env.AdminKey) {
		return true
	}
	return false
}
