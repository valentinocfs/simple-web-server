package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:2222/login", nil)
	recorder := httptest.NewRecorder()

	LoginHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	if bodyString != "Login" {
		t.Errorf("404 Not Found")
	}
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	content := "Login"
	createContent(w, req, content)
}

func createContent(w http.ResponseWriter, r *http.Request, content string) {
	fmt.Fprint(w, content)
}
