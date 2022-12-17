package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hi This is Our First Page</h1>")
}

func contactHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, "<h1>Contact Us </h1> <br> <p> newsun@nw.com </p>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandleFunc)
	r.Get("/contact", contactHandleFunc)
	r.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "Requested Page is not in server", http.StatusNotFound)
	})
	fmt.Println("Starting Server at port 3000")
	http.ListenAndServe(":3000", r)
}
