package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandleFunc(w, r)
	case "/contact":
		contactHandleFunc(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page Not Found")
	}
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hi This is Our First Page</h1>")
}

func contactHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, "<h1>Contact Us </h1> <br> <p> newsun@nw.com </p>")
}

func main() {
	var router Router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	fmt.Println("Starting Server")
	http.ListenAndServe(":3000", router)
}
