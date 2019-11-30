package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chuxinplan/gin-mvc/router"
	"gopkg.in/go-playground/assert.v1"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
