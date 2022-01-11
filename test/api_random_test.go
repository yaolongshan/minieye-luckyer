package test

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestApiGetRandom(t *testing.T) {
//	//for id := 1; id <= 5; id++ {
//	//	for count := 1; count <= 50; count += 5 {
//	//		request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/random?id=%v&count=%v", id, count), nil)
//	//		response, err := client.Do(request)
//	//		if err != nil {
//	//			t.Fail()
//	//		}
//	//		fmt.Println("api: /api/lucky/random", response.StatusCode)
//	//	}
//	//}
//}

func TestApiGetRandom(t *testing.T) {
	//Init()
	//router := r.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/lucky/random?id=1&count=1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
