package test

import (
	"fmt"
	"net/http"
	"testing"
)

var client http.Client

func TestApiRandomGreeting(t *testing.T) {
	for count := 1; count <= 5; count++ {
		request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/greetings/random?&count=%v", count), nil)
		response, err := client.Do(request)
		if err != nil {
			t.Fail()
		}
		fmt.Println("api: /api/greetings/random", response.StatusCode)
	}
}
