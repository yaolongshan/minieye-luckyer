package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestApiGetRandom(t *testing.T) {
	for id := 1; id <= 5; id++ {
		for count := 1; count <= 50; count += 5 {
			request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/random?id=%v&count=%v", id, count), nil)
			response, err := client.Do(request)
			if err != nil {
				t.Fail()
			}
			fmt.Println("api: /api/lucky/random", response.StatusCode)
		}
	}
}
