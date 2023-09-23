package simple_web_server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple Web Go")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/articles", nil)
	recorder := httptest.NewRecorder()

	ArticlesHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodyString)
}
