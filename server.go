package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler404)
	mux.HandleFunc("/login", handlerLogin)
	mux.HandleFunc("/dashboard", handlerDashboard)
	mux.HandleFunc("/member", handlerMember)
	mux.HandleFunc("/report", handlerReport)

	server := http.Server{
		Addr:    "localhost:2222",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func handlerLogin(w http.ResponseWriter, req *http.Request) {
	content := "Login Page"
	createContent(w, req, content)
}

func handlerDashboard(w http.ResponseWriter, req *http.Request) {
	content := "Dashboard Page"
	createContent(w, req, content)
}

func handlerMember(w http.ResponseWriter, req *http.Request) {
	content := "Member Page"
	createContent(w, req, content)
}

func handlerReport(w http.ResponseWriter, req *http.Request) {
	content := "Report Page"
	createContent(w, req, content)
}

func handler404(w http.ResponseWriter, req *http.Request) {
	content := "<h1>404 Not Found</h2>"
	createContent(w, req, content)
}

func createContent(w http.ResponseWriter, r *http.Request, content string) {
	title := "<title>" + r.URL.Path + "</title>"

	html := "<html><head>" + title + "</head>"
	html += "<body>" + content + "<body>"
	html += "</html>"

	fmt.Fprint(w, html)
}
