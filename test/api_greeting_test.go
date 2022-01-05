package test

import (
	r "code/minieye-luckyer/router"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiRandomGreeting(t *testing.T) {
	Init()
	router := r.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/greetings/random?count=1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}


