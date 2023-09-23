package simple_web_server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Simple Web Server")
	})

	mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here's some articles")
	})

	mux.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here's some Blogs")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
