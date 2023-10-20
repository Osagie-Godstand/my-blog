package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", handler)

	r.Get("/about", aboutHandler)

	r.Get("/posts", postsHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/index.tmpl")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/about.tmpl")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/posts.tmpl")
}
