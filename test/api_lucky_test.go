package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestApiGetAllLucky(t *testing.T){
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/list"), nil)
	response, err := client.Do(request)
	if err != nil {
		t.Fail()
	}
	fmt.Println("api: /api/lucky/list",response.StatusCode)
}

func TestApiGetLuckyFile(t *testing.T){
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/api/lucky/file"), nil)
	response, err := client.Do(request)
	if err != nil {
		t.Fail()
	}
	fmt.Println("api: /api/lucky/file",response.StatusCode)
}