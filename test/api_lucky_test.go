package test

import (
	r "code/minieye-luckyer/router"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiGetAllLucky(t *testing.T) {
	//request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/list"), nil)
	//response, err := client.Do(request)
	//if err != nil {
	//	t.Fail()
	//}
	//fmt.Println("api: /api/lucky/list",response.StatusCode)
	Init()
	router := r.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/lucky/list", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestApiGetLuckyFile(t *testing.T) {
	//request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/file"), nil)
	//response, err := client.Do(request)
	//if err != nil {
	//	t.Fail()
	//}
	//fmt.Println("api: /api/lucky/file",response.StatusCode)
	Init()
	router := r.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/lucky/file", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
