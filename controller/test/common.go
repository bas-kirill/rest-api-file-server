package test

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
)

func newReq(method string, path string, body io.Reader, vars map[string]string) *http.Request {
	r := httptest.NewRequest(method, path, body)
	return mux.SetURLVars(r, vars)
}
